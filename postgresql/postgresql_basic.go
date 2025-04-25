package postgresql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"time"

	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/go-observing/notifier"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// BasicPostgresql is the default implementation of the [Postgresql] interface.
type BasicPostgresql struct {
	DatabaseConnector driver.Connector
	isTrace           bool
	logger            logging.Logging
	observerOrigin    string
	observers         subject.Subject
}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var debugOptions = []interface{}{
	&logging.OptionCallerSkip{Value: 5},
}

var traceOptions = []interface{}{
	&logging.OptionCallerSkip{Value: 5},
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
Method GetCurrentWatermark retrieves information about PostgreSQL's transaction IDs.

Input
  - ctx: A context to control lifecycle.

Output
  - The Object Identifier (oid)
  - The age
*/
func (sqlExecutor *BasicPostgresql) GetCurrentWatermark(ctx context.Context) (string, int, error) {
	var (
		oid  string
		age  int
		size string
		err  error
	)

	// Entry tasks.

	if sqlExecutor.isTrace {
		entryTime := time.Now()

		sqlExecutor.traceEntry(1)

		defer func() { sqlExecutor.traceExit(2, oid, age, err, time.Since(entryTime)) }()
	}

	sqlStatement := "SELECT c.oid::regclass, age(c.relfrozenxid), pg_size_pretty(pg_total_relation_size(c.oid)) FROM pg_class c JOIN pg_namespace n on c.relnamespace = n.oid WHERE relkind IN ('r', 't', 'm') AND n.nspname NOT IN ('pg_toast') ORDER BY 2 DESC LIMIT 1;"

	// Open a database connection.

	database := sql.OpenDB(sqlExecutor.DatabaseConnector)
	defer database.Close()

	err = database.PingContext(ctx)
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
			notifier.Notify(ctx, sqlExecutor.observers, sqlExecutor.observerOrigin, ComponentID, 8001, err, details)
		}()
	}

	return oid, age, err
}

/*
Method RegisterObserver adds the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (sqlExecutor *BasicPostgresql) RegisterObserver(ctx context.Context, observer observer.Observer) error {
	var err error

	if sqlExecutor.isTrace {
		entryTime := time.Now()

		sqlExecutor.traceEntry(3, observer.GetObserverID(ctx))

		defer func() { sqlExecutor.traceExit(4, observer.GetObserverID(ctx), err, time.Since(entryTime)) }()
	}

	if sqlExecutor.observers == nil {
		sqlExecutor.observers = &subject.SimpleSubject{}
	}

	err = sqlExecutor.observers.RegisterObserver(ctx, observer)

	if sqlExecutor.observers != nil {
		go func() {
			details := map[string]string{
				"observerID": observer.GetObserverID(ctx),
			}
			notifier.Notify(ctx, sqlExecutor.observers, sqlExecutor.observerOrigin, ComponentID, 8002, err, details)
		}()
	}

	return err
}

/*
Method SetLogLevel sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevel: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (sqlExecutor *BasicPostgresql) SetLogLevel(ctx context.Context, logLevelName string) error {
	var err error

	if sqlExecutor.isTrace {
		entryTime := time.Now()

		sqlExecutor.traceEntry(5, logLevelName)

		defer func() { sqlExecutor.traceExit(6, logLevelName, err, time.Since(entryTime)) }()
	}

	if logging.IsValidLogLevelName(logLevelName) {
		err = sqlExecutor.getLogger().SetLogLevel(logLevelName)
		if err != nil {
			return err
		}

		sqlExecutor.isTrace = (logLevelName == logging.LevelTraceName)
		if sqlExecutor.observers != nil {
			go func() {
				details := map[string]string{
					"logLevelName": logLevelName,
				}
				notifier.Notify(ctx, sqlExecutor.observers, sqlExecutor.observerOrigin, ComponentID, 8003, err, details)
			}()
		}
	} else {
		err = wraperror.Errorf(errPackage, "invalid error level: %s error: %w", logLevelName, errPackage)
	}

	return err
}

/*
Method SetObserverOrigin sets the "origin" value in future Observer messages.

Input
  - ctx: A context to control lifecycle.
  - origin: The value sent in the Observer's "origin" key/value pair.
*/
func (sqlExecutor *BasicPostgresql) SetObserverOrigin(ctx context.Context, origin string) {
	var err error

	// Prolog.

	debugMessageNumber := 0
	traceExitMessageNumber := 69

	if sqlExecutor.getLogger().IsDebug() {
		// If DEBUG, log error exit.
		defer func() {
			if debugMessageNumber > 0 {
				sqlExecutor.debug(debugMessageNumber, err)
			}
		}()

		// If TRACE, Log on entry/exit.

		if sqlExecutor.getLogger().IsTrace() {
			entryTime := time.Now()

			sqlExecutor.traceEntry(60, origin)

			defer func() {
				sqlExecutor.traceExit(traceExitMessageNumber, origin, err, time.Since(entryTime))
			}()
		}

		// If DEBUG, log input parameters. Must be done after establishing DEBUG and TRACE logging.

		asJSON, err := json.Marshal(sqlExecutor) //nolint:musttag
		if err != nil {
			traceExitMessageNumber, debugMessageNumber = 61, 1061

			return
		}

		sqlExecutor.log(1004, sqlExecutor, string(asJSON))
	}

	// Set origin.

	sqlExecutor.observerOrigin = origin

	// Notify observers.

	if sqlExecutor.observers != nil {
		go func() {
			details := map[string]string{
				"origin": origin,
			}
			notifier.Notify(ctx, sqlExecutor.observers, sqlExecutor.observerOrigin, ComponentID, 8005, err, details)
		}()
	}
}

/*
Method UnregisterObserver removes the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (sqlExecutor *BasicPostgresql) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
	var err error

	if sqlExecutor.isTrace {
		entryTime := time.Now()

		sqlExecutor.traceEntry(7, observer.GetObserverID(ctx))

		defer func() { sqlExecutor.traceExit(8, observer.GetObserverID(ctx), err, time.Since(entryTime)) }()
	}

	if sqlExecutor.observers != nil {
		// Tricky code:
		// client.notify is called synchronously before client.observers is set to nil.
		// In client.notify, each observer will get notified in a goroutine.
		// Then client.observers may be set to nil, but observer goroutines will be OK.
		details := map[string]string{
			"observerID": observer.GetObserverID(ctx),
		}
		notifier.Notify(ctx, sqlExecutor.observers, sqlExecutor.observerOrigin, ComponentID, 8004, err, details)
	}

	err = sqlExecutor.observers.UnregisterObserver(ctx, observer)

	if !sqlExecutor.observers.HasObservers(ctx) {
		sqlExecutor.observers = nil
	}

	return err
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// Get the Logger singleton.
func (sqlExecutor *BasicPostgresql) getLogger() logging.Logging {
	var err error

	if sqlExecutor.logger == nil {
		options := []interface{}{
			&logging.OptionCallerSkip{Value: 4},
		}

		sqlExecutor.logger, err = logging.NewSenzingLogger(ComponentID, IDMessages, options...)
		if err != nil {
			panic(err)
		}
	}

	return sqlExecutor.logger
}

// Log message.
func (sqlExecutor *BasicPostgresql) log(messageNumber int, details ...interface{}) {
	sqlExecutor.getLogger().Log(messageNumber, details...)
}

// Debug.
func (sqlExecutor *BasicPostgresql) debug(messageNumber int, details ...interface{}) {
	details = append(details, debugOptions...)
	sqlExecutor.getLogger().Log(messageNumber, details...)
}

// Trace method entry.
func (sqlExecutor *BasicPostgresql) traceEntry(messageNumber int, details ...interface{}) {
	details = append(details, traceOptions...)
	sqlExecutor.getLogger().Log(messageNumber, details...)
}

// Trace method exit.
func (sqlExecutor *BasicPostgresql) traceExit(messageNumber int, details ...interface{}) {
	details = append(details, traceOptions...)
	sqlExecutor.getLogger().Log(messageNumber, details...)
}
