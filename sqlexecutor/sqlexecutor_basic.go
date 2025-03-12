package sqlexecutor

import (
	"bufio"
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/go-observing/notifier"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// BasicSQLExecutor is the default implementation of the [SQLExecutor] interface.
type BasicSQLExecutor struct {
	DatabaseConnector driver.Connector `json:"databaseConnector,omitempty"`
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
Method ProcessFileName is a convenience method for calling [BasicSQLExecutor.ProcessScanner] using a filename.

Input
  - ctx: A context to control lifecycle.
  - filename: A fully qualified path to a file of SQL statements.
*/
func (sqlExecutor *BasicSQLExecutor) ProcessFileName(ctx context.Context, filename string) error {

	var err error

	// Entry tasks.

	if sqlExecutor.isTrace {
		entryTime := time.Now()
		sqlExecutor.traceEntry(1, filename)
		defer func() { sqlExecutor.traceExit(2, filename, err, time.Since(entryTime)) }()
	}

	// Process file.

	filename = filepath.Clean(filename)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			sqlExecutor.log(9999, file, err)
		}
	}()
	err = sqlExecutor.ProcessScanner(ctx, bufio.NewScanner(file))
	if err != nil {
		return err
	}

	// Exit tasks.

	if sqlExecutor.observers != nil {
		go func() {
			details := map[string]string{
				"filename": filename,
			}
			notifier.Notify(ctx, sqlExecutor.observers, sqlExecutor.observerOrigin, ComponentID, 8001, err, details)
		}()
	}
	return err
}

/*
Method ProcessScanner does a database call for each line scanned.

Input
  - ctx: A context to control lifecycle.
  - scanner: SQL statements to be processed.
*/
func (sqlExecutor *BasicSQLExecutor) ProcessScanner(ctx context.Context, scanner *bufio.Scanner) error {

	var (
		err         error
		scanLine    = 0
		scanFailure = 0
	)
	// Entry tasks.

	if sqlExecutor.isTrace {
		entryTime := time.Now()
		sqlExecutor.traceEntry(3)
		defer func() { sqlExecutor.traceExit(4, scanLine, scanFailure, err, time.Since(entryTime)) }()
	}

	// Open a database connection.

	database := sql.OpenDB(sqlExecutor.DatabaseConnector)
	// defer database.Close()

	err = database.PingContext(ctx)
	if err != nil {
		return err
	}

	// Process each scanned line.

	for scanner.Scan() {
		scanLine++
		sqlText := scanner.Text()
		result, err := database.ExecContext(ctx, sqlText)
		if err != nil {
			scanFailure++
			sqlExecutor.log(3001, scanFailure, scanLine, result, err)
		}
		if sqlExecutor.observers != nil {
			go func() {
				details := map[string]string{
					"line": strconv.Itoa(scanLine),
					"SQL":  sqlText,
				}
				notifier.Notify(ctx, sqlExecutor.observers, sqlExecutor.observerOrigin, ComponentID, 8002, err, details)
			}()
		}
	}
	err = scanner.Err()

	// >>>>> FIXME:

	// sqlRows, err := database.Query("SELECT name FROM sqlite_master WHERE type='table';")
	// if err != nil {
	// 	return err
	// }
	// defer sqlRows.Close()

	// var name string
	// for sqlRows.Next() {
	// 	err := sqlRows.Scan(&name)
	// 	if err != nil {
	// 		fmt.Printf(">>>>> error: %v\n", err)
	// 	}
	// 	fmt.Printf(">>>>> name: %s\n", name)
	// }

	// Exit tasks.

	if sqlExecutor.observers != nil {
		go func() {
			details := map[string]string{
				"lines":    strconv.Itoa(scanLine),
				"failures": strconv.Itoa(scanFailure),
			}
			notifier.Notify(ctx, sqlExecutor.observers, sqlExecutor.observerOrigin, ComponentID, 8003, err, details)
		}()
	}

	// Determine error level to log.

	messageNumber := 2001 // INFO message.
	if scanFailure > 0 {
		messageNumber = 3002 // WARN message.
	}
	sqlExecutor.log(messageNumber, scanLine, scanFailure)

	return err
}

/*
Method RegisterObserver adds the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (sqlExecutor *BasicSQLExecutor) RegisterObserver(ctx context.Context, observer observer.Observer) error {
	var err error

	if sqlExecutor.isTrace {
		entryTime := time.Now()
		sqlExecutor.traceEntry(5, observer.GetObserverID(ctx))
		defer func() { sqlExecutor.traceExit(6, observer.GetObserverID(ctx), err, time.Since(entryTime)) }()
	}
	if sqlExecutor.observers == nil {
		sqlExecutor.observers = &subject.SimpleSubject{}
	}

	// Register observer with sqlExecutor.

	err = sqlExecutor.observers.RegisterObserver(ctx, observer)

	// Notify observers.

	go func() {
		details := map[string]string{
			"observerID": observer.GetObserverID(ctx),
		}
		notifier.Notify(ctx, sqlExecutor.observers, sqlExecutor.observerOrigin, ComponentID, 8004, err, details)
	}()

	// Epilog.

	return err
}

/*
Method SetLogLevel sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevelName: The desired log level as string: "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL" or "PANIC".
*/
func (sqlExecutor *BasicSQLExecutor) SetLogLevel(ctx context.Context, logLevelName string) error {
	var err error

	if sqlExecutor.isTrace {
		entryTime := time.Now()
		sqlExecutor.traceEntry(7, logLevelName)
		defer func() { sqlExecutor.traceExit(8, logLevelName, err, time.Since(entryTime)) }()
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
				notifier.Notify(ctx, sqlExecutor.observers, sqlExecutor.observerOrigin, ComponentID, 8005, err, details)
			}()
		}
	} else {
		err = fmt.Errorf("invalid error level: %s", logLevelName)
	}

	return err
}

/*
Method SetObserverOrigin sets the "origin" value in future Observer messages.

Input
  - ctx: A context to control lifecycle.
  - origin: The value sent in the Observer's "origin" key/value pair.
*/
func (sqlExecutor *BasicSQLExecutor) SetObserverOrigin(ctx context.Context, origin string) {
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

		asJSON, err := json.Marshal(sqlExecutor)
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
func (sqlExecutor *BasicSQLExecutor) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
	var err error

	if sqlExecutor.isTrace {
		entryTime := time.Now()
		sqlExecutor.traceEntry(9, observer.GetObserverID(ctx))
		defer func() { sqlExecutor.traceExit(10, observer.GetObserverID(ctx), err, time.Since(entryTime)) }()
	}

	// Remove observer from this service.

	if sqlExecutor.observers != nil {
		err = sqlExecutor.observers.UnregisterObserver(ctx, observer)
		if err != nil {
			return err
		}

		// Tricky code:
		// client.notify is called synchronously before client.observers is set to nil.
		// In client.notify, each observer will get notified in a goroutine.
		// Then client.observers may be set to nil, but observer goroutines will be OK.
		details := map[string]string{
			"observerID": observer.GetObserverID(ctx),
		}
		notifier.Notify(ctx, sqlExecutor.observers, sqlExecutor.observerOrigin, ComponentID, 8006, err, details)

		if !sqlExecutor.observers.HasObservers(ctx) {
			sqlExecutor.observers = nil
		}
	}

	// Epilog.

	return err
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// --- Logging ----------------------------------------------------------------

// Get the Logger singleton.
func (sqlExecutor *BasicSQLExecutor) getLogger() logging.Logging {
	var err error
	if sqlExecutor.logger == nil {
		options := []interface{}{
			logging.OptionCallerSkip{Value: 4},
			logging.OptionMessageFields{Value: []string{"id", "text"}},
		}
		sqlExecutor.logger, err = logging.NewSenzingLogger(ComponentID, IDMessages, options...)
		if err != nil {
			panic(err)
		}
	}
	return sqlExecutor.logger
}

// Log message.
func (sqlExecutor *BasicSQLExecutor) log(messageNumber int, details ...interface{}) {
	sqlExecutor.getLogger().Log(messageNumber, details...)
}

// Debug.
func (sqlExecutor *BasicSQLExecutor) debug(messageNumber int, details ...interface{}) {
	details = append(details, debugOptions...)
	sqlExecutor.getLogger().Log(messageNumber, details...)
}

// Trace method entry.
func (sqlExecutor *BasicSQLExecutor) traceEntry(messageNumber int, details ...interface{}) {
	details = append(details, traceOptions...)
	sqlExecutor.getLogger().Log(messageNumber, details...)
}

// Trace method exit.
func (sqlExecutor *BasicSQLExecutor) traceExit(messageNumber int, details ...interface{}) {
	details = append(details, traceOptions...)
	sqlExecutor.getLogger().Log(messageNumber, details...)
}
