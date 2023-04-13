package sqlexecutor

import (
	"bufio"
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"os"
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

// SqlExecutorImpl is the default implementation of the SqlExecutor interface.
type SqlExecutorImpl struct {
	DatabaseConnector driver.Connector
	isTrace           bool
	logger            logging.LoggingInterface
	observers         subject.Subject
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// Get the Logger singleton.
func (sqlExecutor *SqlExecutorImpl) getLogger() logging.LoggingInterface {
	var err error = nil
	if sqlExecutor.logger == nil {
		options := []interface{}{
			&logging.OptionCallerSkip{Value: 4},
		}
		sqlExecutor.logger, err = logging.NewSenzingToolsLogger(ProductId, IdMessages, options...)
		if err != nil {
			panic(err)
		}
	}
	return sqlExecutor.logger
}

// Log message.
func (sqlExecutor *SqlExecutorImpl) log(messageNumber int, details ...interface{}) {
	sqlExecutor.getLogger().Log(messageNumber, details...)
}

// Trace method entry.
func (sqlExecutor *SqlExecutorImpl) traceEntry(messageNumber int, details ...interface{}) {
	sqlExecutor.getLogger().Log(messageNumber, details...)
}

// Trace method exit.
func (sqlExecutor *SqlExecutorImpl) traceExit(messageNumber int, details ...interface{}) {
	sqlExecutor.getLogger().Log(messageNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The ProcessFileName is a convenience method for calling method ProcessScanner using a filename.

Input
  - ctx: A context to control lifecycle.
  - filename: A fully qualified path to a file of SQL statements.
*/
func (sqlExecutor *SqlExecutorImpl) ProcessFileName(ctx context.Context, filename string) error {

	// Entry tasks.

	if sqlExecutor.isTrace {
		sqlExecutor.traceEntry(1, filename)
	}
	entryTime := time.Now()

	// Process file.

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
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
			notifier.Notify(ctx, sqlExecutor.observers, ProductId, 8001, err, details)
		}()
	}
	if sqlExecutor.isTrace {
		defer sqlExecutor.traceExit(2, filename, err, time.Since(entryTime))
	}
	return err
}

/*
The ProcessScanner does a database call for each line scanned.

Input
  - ctx: A context to control lifecycle.
  - scanner: SQL statements to be processed.
*/
func (sqlExecutor *SqlExecutorImpl) ProcessScanner(ctx context.Context, scanner *bufio.Scanner) error {

	// Entry tasks.

	if sqlExecutor.isTrace {
		sqlExecutor.traceEntry(3)
	}
	entryTime := time.Now()

	// Open a database connection.

	database := sql.OpenDB(sqlExecutor.DatabaseConnector)
	defer database.Close()

	err := database.PingContext(ctx)
	if err != nil {
		return err
	}

	// Process each scanned line.

	scanLine := 0
	scanFailure := 0
	for scanner.Scan() {
		scanLine += 1
		sqlText := scanner.Text()
		result, err := database.ExecContext(ctx, sqlText)
		if err != nil {
			scanFailure += 1
			sqlExecutor.log(3001, scanFailure, scanLine, result, err)
		}
		if sqlExecutor.observers != nil {
			go func() {
				details := map[string]string{
					"line": strconv.Itoa(scanLine),
					"SQL":  sqlText,
				}
				notifier.Notify(ctx, sqlExecutor.observers, ProductId, 8002, err, details)
			}()
		}
	}
	err = scanner.Err()

	// Exit tasks.

	if sqlExecutor.observers != nil {
		go func() {
			details := map[string]string{
				"lines":    strconv.Itoa(scanLine),
				"failures": strconv.Itoa(scanFailure),
			}
			notifier.Notify(ctx, sqlExecutor.observers, ProductId, 8003, err, details)
		}()
	}

	// Determine error level to log.

	messageNumber := 2001
	if scanFailure > 0 {
		messageNumber = 3002
	}
	sqlExecutor.log(messageNumber, scanLine, scanFailure)

	if sqlExecutor.isTrace {
		defer sqlExecutor.traceExit(4, scanLine, scanFailure, err, time.Since(entryTime))
	}
	return err
}

/*
The RegisterObserver method adds the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (sqlExecutor *SqlExecutorImpl) RegisterObserver(ctx context.Context, observer observer.Observer) error {
	if sqlExecutor.isTrace {
		sqlExecutor.traceEntry(5, observer.GetObserverId(ctx))
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
			notifier.Notify(ctx, sqlExecutor.observers, ProductId, 8004, err, details)
		}()
	}
	if sqlExecutor.isTrace {
		defer sqlExecutor.traceExit(6, observer.GetObserverId(ctx), err, time.Since(entryTime))
	}
	return err
}

/*
The SetLogLevel method sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevelName: The desired log level as string: "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL" or "PANIC".
*/
func (sqlExecutor *SqlExecutorImpl) SetLogLevel(ctx context.Context, logLevelName string) error {
	if sqlExecutor.isTrace {
		sqlExecutor.traceEntry(7, logLevelName)
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
				notifier.Notify(ctx, sqlExecutor.observers, ProductId, 8005, err, details)
			}()
		}
	} else {
		err = fmt.Errorf("invalid error level: %s", logLevelName)
	}
	if sqlExecutor.isTrace {
		defer sqlExecutor.traceExit(8, logLevelName, err, time.Since(entryTime))
	}
	return err
}

/*
The UnregisterObserver method removes the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (sqlExecutor *SqlExecutorImpl) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
	if sqlExecutor.isTrace {
		sqlExecutor.traceEntry(9, observer.GetObserverId(ctx))
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
		notifier.Notify(ctx, sqlExecutor.observers, ProductId, 8006, err, details)
	}
	err = sqlExecutor.observers.UnregisterObserver(ctx, observer)
	if !sqlExecutor.observers.HasObservers(ctx) {
		sqlExecutor.observers = nil
	}
	if sqlExecutor.isTrace {
		defer sqlExecutor.traceExit(10, observer.GetObserverId(ctx), err, time.Since(entryTime))
	}
	return err
}
