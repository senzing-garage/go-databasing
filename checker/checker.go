package checker

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
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

// CheckerImpl is the default implementation of the SqlExecutor interface.
type CheckerImpl struct {
	DatabaseConnector driver.Connector
	isTrace           bool
	logger            logging.LoggingInterface
	observerOrigin    string
	observers         subject.Subject
}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var debugOptions []interface{} = []interface{}{
	&logging.OptionCallerSkip{Value: 5},
}

var traceOptions []interface{} = []interface{}{
	&logging.OptionCallerSkip{Value: 5},
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// Get the Logger singleton.
func (schemaChecker *CheckerImpl) getLogger() logging.LoggingInterface {
	var err error = nil
	if schemaChecker.logger == nil {
		options := []interface{}{
			&logging.OptionCallerSkip{Value: 4},
		}
		schemaChecker.logger, err = logging.NewSenzingToolsLogger(ComponentId, IdMessages, options...)
		if err != nil {
			panic(err)
		}
	}
	return schemaChecker.logger
}

// Log message.
func (schemaChecker *CheckerImpl) log(messageNumber int, details ...interface{}) {
	schemaChecker.getLogger().Log(messageNumber, details...)
}

// Debug.
func (schemaChecker *CheckerImpl) debug(messageNumber int, details ...interface{}) {
	details = append(details, debugOptions...)
	schemaChecker.getLogger().Log(messageNumber, details...)
}

// Trace method entry.
func (schemaChecker *CheckerImpl) traceEntry(messageNumber int, details ...interface{}) {
	details = append(details, traceOptions...)
	schemaChecker.getLogger().Log(messageNumber, details...)
}

// Trace method exit.
func (schemaChecker *CheckerImpl) traceExit(messageNumber int, details ...interface{}) {
	details = append(details, traceOptions...)
	schemaChecker.getLogger().Log(messageNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The IsInstalled verifies that the Senzing schema has been installed.

Input
  - ctx: A context to control lifecycle.
*/
func (schemaChecker *CheckerImpl) IsSchemaInstalled(ctx context.Context) (bool, error) {
	var (
		count int
	)

	// Entry tasks.

	if schemaChecker.isTrace {
		schemaChecker.traceEntry(1)
	}
	entryTime := time.Now()
	sqlStatement := "SELECT count(*) from DSRC_RECORD;"

	// Open a database connection.

	database := sql.OpenDB(schemaChecker.DatabaseConnector)
	defer database.Close()
	err := database.PingContext(ctx)
	if err != nil {
		return false, err
	}

	// Get the Row.

	row := database.QueryRowContext(ctx, sqlStatement)
	err = row.Scan(&count)
	if err != nil {
		return false, err
	}

	// Exit tasks.

	if schemaChecker.observers != nil {
		go func() {
			details := map[string]string{
				"count": strconv.Itoa(count),
			}
			notifier.Notify(ctx, schemaChecker.observers, schemaChecker.observerOrigin, ComponentId, 8001, err, details)
		}()
	}
	if schemaChecker.isTrace {
		defer schemaChecker.traceExit(2, count, err, time.Since(entryTime))
	}
	return true, err
}

/*
The IsInstalled verifies that the Senzing schema has been installed.

Input
  - ctx: A context to control lifecycle.
*/
func (schemaChecker *CheckerImpl) RecordCount(ctx context.Context) (int64, error) {
	var (
		count int64
	)

	// Entry tasks.

	if schemaChecker.isTrace {
		schemaChecker.traceEntry(9)
	}
	entryTime := time.Now()
	sqlStatement := "SELECT count(*) from DSRC_RECORD;"

	// Open a database connection.

	database := sql.OpenDB(schemaChecker.DatabaseConnector)
	defer database.Close()
	err := database.PingContext(ctx)
	if err != nil {
		return 0, err
	}

	// Get the Row.

	row := database.QueryRowContext(ctx, sqlStatement)
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	// Exit tasks.

	if schemaChecker.observers != nil {
		go func() {
			details := map[string]string{
				"count": strconv.FormatInt(count, 10),
			}
			notifier.Notify(ctx, schemaChecker.observers, schemaChecker.observerOrigin, ComponentId, 8005, err, details)
		}()
	}
	if schemaChecker.isTrace {
		defer schemaChecker.traceExit(10, count, err, time.Since(entryTime))
	}
	return count, err
}

/*
The RegisterObserver method adds the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (schemaChecker *CheckerImpl) RegisterObserver(ctx context.Context, observer observer.Observer) error {
	if schemaChecker.isTrace {
		schemaChecker.traceEntry(3, observer.GetObserverId(ctx))
	}
	entryTime := time.Now()
	if schemaChecker.observers == nil {
		schemaChecker.observers = &subject.SubjectImpl{}
	}
	err := schemaChecker.observers.RegisterObserver(ctx, observer)
	if schemaChecker.observers != nil {
		go func() {
			details := map[string]string{
				"observerID": observer.GetObserverId(ctx),
			}
			notifier.Notify(ctx, schemaChecker.observers, schemaChecker.observerOrigin, ComponentId, 8002, err, details)
		}()
	}
	if schemaChecker.isTrace {
		defer schemaChecker.traceExit(4, observer.GetObserverId(ctx), err, time.Since(entryTime))
	}
	return err
}

/*
The SetLogLevel method sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevel: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (schemaChecker *CheckerImpl) SetLogLevel(ctx context.Context, logLevelName string) error {
	if schemaChecker.isTrace {
		schemaChecker.traceEntry(5, logLevelName)
	}
	entryTime := time.Now()
	var err error = nil
	if logging.IsValidLogLevelName(logLevelName) {
		err = schemaChecker.getLogger().SetLogLevel(logLevelName)
		if err != nil {
			return err
		}
		schemaChecker.isTrace = (logLevelName == logging.LevelTraceName)
		if schemaChecker.observers != nil {
			go func() {
				details := map[string]string{
					"logLevelName": logLevelName,
				}
				notifier.Notify(ctx, schemaChecker.observers, schemaChecker.observerOrigin, ComponentId, 8003, err, details)
			}()
		}
	} else {
		err = fmt.Errorf("invalid error level: %s", logLevelName)
	}
	if schemaChecker.isTrace {
		defer schemaChecker.traceExit(6, logLevelName, err, time.Since(entryTime))
	}
	return err
}

/*
The SetObserverOrigin method sets the "origin" value in future Observer messages.

Input
  - ctx: A context to control lifecycle.
  - origin: The value sent in the Observer's "origin" key/value pair.
*/
func (schemaChecker *CheckerImpl) SetObserverOrigin(ctx context.Context, origin string) {
	var err error = nil

	// Prolog.

	debugMessageNumber := 0
	traceExitMessageNumber := 69
	if schemaChecker.getLogger().IsDebug() {

		// If DEBUG, log error exit.

		defer func() {
			if debugMessageNumber > 0 {
				schemaChecker.debug(debugMessageNumber, err)
			}
		}()

		// If TRACE, Log on entry/exit.

		if schemaChecker.getLogger().IsTrace() {
			entryTime := time.Now()
			schemaChecker.traceEntry(60, origin)
			defer func() {
				schemaChecker.traceExit(traceExitMessageNumber, origin, err, time.Since(entryTime))
			}()
		}

		// If DEBUG, log input parameters. Must be done after establishing DEBUG and TRACE logging.

		asJson, err := json.Marshal(schemaChecker)
		if err != nil {
			traceExitMessageNumber, debugMessageNumber = 61, 1061
			return
		}
		schemaChecker.log(1004, schemaChecker, string(asJson))
	}

	// Set origin.

	schemaChecker.observerOrigin = origin

	// Notify observers.

	if schemaChecker.observers != nil {
		go func() {
			details := map[string]string{
				"origin": origin,
			}
			notifier.Notify(ctx, schemaChecker.observers, schemaChecker.observerOrigin, ComponentId, 8005, err, details)
		}()
	}

}

/*
The UnregisterObserver method removes the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (schemaChecker *CheckerImpl) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
	if schemaChecker.isTrace {
		schemaChecker.traceEntry(7, observer.GetObserverId(ctx))
	}
	entryTime := time.Now()
	var err error = nil
	if schemaChecker.observers != nil {
		// Tricky code:
		// client.notify is called synchronously before client.observers is set to nil.
		// In client.notify, each observer will get notified in a goroutine.
		// Then client.observers may be set to nil, but observer goroutines will be OK.
		details := map[string]string{
			"observerID": observer.GetObserverId(ctx),
		}
		notifier.Notify(ctx, schemaChecker.observers, schemaChecker.observerOrigin, ComponentId, 8004, err, details)
	}
	err = schemaChecker.observers.UnregisterObserver(ctx, observer)
	if !schemaChecker.observers.HasObservers(ctx) {
		schemaChecker.observers = nil
	}
	if schemaChecker.isTrace {
		defer schemaChecker.traceExit(8, observer.GetObserverId(ctx), err, time.Since(entryTime))
	}
	return err
}
