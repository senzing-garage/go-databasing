package postgresql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"

	"github.com/senzing/go-logging/logging"
	"github.com/senzing/go-observing/notifier"
	"github.com/senzing/go-observing/observer"
	"github.com/senzing/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// PostgresqlImpl is the default implementation of the SqlExecutor interface.
type PostgresqlImpl struct {
	DatabaseConnector driver.Connector
	isTrace           bool
	logger            logging.LoggingInterface
	observers         subject.Subject
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// Get the Logger singleton.
func (sqlExecutor *PostgresqlImpl) getLogger() logging.LoggingInterface {
	var err error = nil
	if sqlExecutor.logger == nil {
		sqlExecutor.logger, err = logging.NewSenzingToolsLogger(ProductId, IdMessages)
		if err != nil {
			panic(err)
		}
	}
	return sqlExecutor.logger
}

// Log message.
func (sqlExecutor *PostgresqlImpl) log(messageNumber int, details ...interface{}) {
	sqlExecutor.getLogger().Log(messageNumber, details...)
}

// Trace method entry.
func (sqlExecutor *PostgresqlImpl) traceEntry(errorNumber int, details ...interface{}) {
	sqlExecutor.log(errorNumber, details...)
}

// Trace method exit.
func (sqlExecutor *PostgresqlImpl) traceExit(errorNumber int, details ...interface{}) {
	sqlExecutor.log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The GetCurrentWatermark does a database call for each line scanned.

Input
  - ctx: A context to control lifecycle.
*/
func (sqlExecutor *PostgresqlImpl) GetCurrentWatermark(ctx context.Context) (string, int, error) {
	var (
		oid  string
		age  int
		size string
	)

	// Entry tasks.

	if sqlExecutor.isTrace {
		sqlExecutor.traceEntry(1)
	}
	entryTime := time.Now()
	sqlStatement := "SELECT c.oid::regclass, age(c.relfrozenxid), pg_size_pretty(pg_total_relation_size(c.oid)) FROM pg_class c JOIN pg_namespace n on c.relnamespace = n.oid WHERE relkind IN ('r', 't', 'm') AND n.nspname NOT IN ('pg_toast') ORDER BY 2 DESC LIMIT 1;"

	// Open a database connection.

	database := sql.OpenDB(sqlExecutor.DatabaseConnector)
	defer database.Close()
	err := database.PingContext(ctx)
	if err != nil {
		return "", 0, err
	}

	// Get the Row.

	row := database.QueryRowContext(ctx, sqlStatement)
	err = row.Scan(&oid, &age, &size)
	if err != nil {
		return "", 0, err
	}

	// Exit tasks.

	if sqlExecutor.observers != nil {
		go func() {
			details := map[string]string{
				"oid": oid,
				"age": strconv.Itoa(age),
			}
			notifier.Notify(ctx, sqlExecutor.observers, ProductId, 8001, err, details)
		}()
	}
	if sqlExecutor.isTrace {
		defer sqlExecutor.traceExit(2, oid, age, err, time.Since(entryTime))
	}
	return oid, age, err
}

/*
The RegisterObserver method adds the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (sqlExecutor *PostgresqlImpl) RegisterObserver(ctx context.Context, observer observer.Observer) error {
	if sqlExecutor.isTrace {
		sqlExecutor.traceEntry(3, observer.GetObserverId(ctx))
	}
	entryTime := time.Now()
	if sqlExecutor.observers == nil {
		sqlExecutor.observers = &subject.SubjectImpl{}
	}
	err := sqlExecutor.observers.RegisterObserver(ctx, observer)
	if sqlExecutor.observers != nil {
		go func() {
			details := map[string]string{
				"observerID": observer.GetObserverId(ctx),
			}
			notifier.Notify(ctx, sqlExecutor.observers, ProductId, 8002, err, details)
		}()
	}
	if sqlExecutor.isTrace {
		defer sqlExecutor.traceExit(4, observer.GetObserverId(ctx), err, time.Since(entryTime))
	}
	return err
}

/*
The SetLogLevel method sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevel: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (sqlExecutor *PostgresqlImpl) SetLogLevel(ctx context.Context, logLevelName string) error {
	if sqlExecutor.isTrace {
		sqlExecutor.traceEntry(5, logLevelName)
	}
	entryTime := time.Now()
	var err error = nil
	if logging.IsValidLogLevelName(logLevelName) {
		sqlExecutor.getLogger().SetLogLevel(logLevelName)
		sqlExecutor.isTrace = (logLevelName == logging.LevelTraceName)
		if sqlExecutor.observers != nil {
			go func() {
				details := map[string]string{
					"logLevelName": logLevelName,
				}
				notifier.Notify(ctx, sqlExecutor.observers, ProductId, 8003, err, details)
			}()
		}
	} else {
		err = fmt.Errorf("invalid error level: %s", logLevelName)
	}
	if sqlExecutor.isTrace {
		defer sqlExecutor.traceExit(6, logLevelName, err, time.Since(entryTime))
	}
	return err
}

/*
The UnregisterObserver method removes the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (sqlExecutor *PostgresqlImpl) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
	if sqlExecutor.isTrace {
		sqlExecutor.traceEntry(7, observer.GetObserverId(ctx))
	}
	entryTime := time.Now()
	var err error = nil
	if sqlExecutor.observers != nil {
		// Tricky code:
		// client.notify is called synchronously before client.observers is set to nil.
		// In client.notify, each observer will get notified in a goroutine.
		// Then client.observers may be set to nil, but observer goroutines will be OK.
		details := map[string]string{
			"observerID": observer.GetObserverId(ctx),
		}
		notifier.Notify(ctx, sqlExecutor.observers, ProductId, 8004, err, details)
	}
	err = sqlExecutor.observers.UnregisterObserver(ctx, observer)
	if !sqlExecutor.observers.HasObservers(ctx) {
		sqlExecutor.observers = nil
	}
	if sqlExecutor.isTrace {
		defer sqlExecutor.traceExit(8, observer.GetObserverId(ctx), err, time.Since(entryTime))
	}
	return err
}
