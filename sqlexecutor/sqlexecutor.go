package sqlexecutor

import (
	"bufio"
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagelogger"
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
	logger            messagelogger.MessageLoggerInterface
	LogLevel          logger.Level
	observers         subject.Subject
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// Get the Logger singleton.
func (sqlexecutor *SqlExecutorImpl) getLogger() messagelogger.MessageLoggerInterface {
	if sqlexecutor.logger == nil {
		sqlexecutor.logger, _ = messagelogger.NewSenzingApiLogger(ProductId, IdMessages, IdStatuses, messagelogger.LevelInfo)
	}
	return sqlexecutor.logger
}

// Notify registered observers.
func (sqlexecutor *SqlExecutorImpl) notify(ctx context.Context, messageId int, err error, details map[string]string) {
	now := time.Now()
	details["subjectId"] = strconv.Itoa(ProductId)
	details["messageId"] = strconv.Itoa(messageId)
	details["messageTime"] = strconv.FormatInt(now.UnixNano(), 10)
	if err != nil {
		details["error"] = err.Error()
	}
	message, err := json.Marshal(details)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		sqlexecutor.observers.NotifyObservers(ctx, string(message))
	}
}

// Trace method entry.
func (sqlexecutor *SqlExecutorImpl) traceEntry(errorNumber int, details ...interface{}) {
	sqlexecutor.getLogger().Log(errorNumber, details...)
}

// Trace method exit.
func (sqlexecutor *SqlExecutorImpl) traceExit(errorNumber int, details ...interface{}) {
	sqlexecutor.getLogger().Log(errorNumber, details...)
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
func (sqlexecutor *SqlExecutorImpl) ProcessFileName(ctx context.Context, filename string) error {

	// Entry tasks.

	if sqlexecutor.isTrace {
		sqlexecutor.traceEntry(1, filename)
	}
	var err error = nil
	entryTime := time.Now()

	// Process file.

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	sqlexecutor.ProcessScanner(ctx, bufio.NewScanner(file))

	// Exit tasks.

	if sqlexecutor.observers != nil {
		go func() {
			details := map[string]string{
				"filename": filename,
			}
			sqlexecutor.notify(ctx, 8001, err, details)
		}()
	}
	if sqlexecutor.isTrace {
		defer sqlexecutor.traceExit(2, filename, err, time.Since(entryTime))
	}
	return err
}

/*
The ProcessScanner does a database call for each line scanned.

Input
  - ctx: A context to control lifecycle.
  - scanner: SQL statements to be processed.
*/
func (sqlexecutor *SqlExecutorImpl) ProcessScanner(ctx context.Context, scanner *bufio.Scanner) {

	// Entry tasks.

	if sqlexecutor.isTrace {
		sqlexecutor.traceEntry(3)
	}
	var err error = nil
	entryTime := time.Now()

	// Open a database connection.

	database := sql.OpenDB(sqlexecutor.DatabaseConnector)
	defer database.Close()

	// Process each scanned line.

	for scanner.Scan() {
		sqlText := scanner.Text()
		result, err := database.ExecContext(ctx, sqlText)
		if err != nil {
			sqlexecutor.getLogger().Log(3001, result, err)
		}
		if sqlexecutor.observers != nil {
			go func() {
				details := map[string]string{
					"SQL": sqlText,
				}
				sqlexecutor.notify(ctx, 8002, err, details)
			}()
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Exit tasks.

	if sqlexecutor.observers != nil {
		go func() {
			details := map[string]string{}
			sqlexecutor.notify(ctx, 8003, err, details)
		}()
	}
	if sqlexecutor.isTrace {
		defer sqlexecutor.traceExit(4, err, time.Since(entryTime))
	}
}

/*
The RegisterObserver method adds the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (sqlexecutor *SqlExecutorImpl) RegisterObserver(ctx context.Context, observer observer.Observer) error {
	if sqlexecutor.isTrace {
		sqlexecutor.traceEntry(5, observer.GetObserverId(ctx))
	}
	entryTime := time.Now()
	if sqlexecutor.observers == nil {
		sqlexecutor.observers = &subject.SubjectImpl{}
	}
	err := sqlexecutor.observers.RegisterObserver(ctx, observer)
	if sqlexecutor.observers != nil {
		go func() {
			details := map[string]string{
				"observerID": observer.GetObserverId(ctx),
			}
			sqlexecutor.notify(ctx, 8004, err, details)
		}()
	}
	if sqlexecutor.isTrace {
		defer sqlexecutor.traceExit(6, observer.GetObserverId(ctx), err, time.Since(entryTime))
	}
	return err
}

/*
The SetLogLevel method sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevel: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (sqlexecutor *SqlExecutorImpl) SetLogLevel(ctx context.Context, logLevel logger.Level) error {
	if sqlexecutor.isTrace {
		sqlexecutor.traceEntry(7, logLevel)
	}
	entryTime := time.Now()
	var err error = nil
	sqlexecutor.getLogger().SetLogLevel(messagelogger.Level(logLevel))
	sqlexecutor.isTrace = (sqlexecutor.getLogger().GetLogLevel() == messagelogger.LevelTrace)
	if sqlexecutor.observers != nil {
		go func() {
			details := map[string]string{
				"logLevel": logger.LevelToTextMap[logLevel],
			}
			sqlexecutor.notify(ctx, 8005, err, details)
		}()
	}
	if sqlexecutor.isTrace {
		defer sqlexecutor.traceExit(8, logLevel, err, time.Since(entryTime))
	}
	return err
}

/*
The UnregisterObserver method removes the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (sqlexecutor *SqlExecutorImpl) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
	if sqlexecutor.isTrace {
		sqlexecutor.traceEntry(9, observer.GetObserverId(ctx))
	}
	entryTime := time.Now()
	var err error = nil
	if sqlexecutor.observers != nil {
		// Tricky code:
		// client.notify is called synchronously before client.observers is set to nil.
		// In client.notify, each observer will get notified in a goroutine.
		// Then client.observers may be set to nil, but observer goroutines will be OK.
		details := map[string]string{
			"observerID": observer.GetObserverId(ctx),
		}
		sqlexecutor.notify(ctx, 8006, err, details)
	}
	err = sqlexecutor.observers.UnregisterObserver(ctx, observer)
	if !sqlexecutor.observers.HasObservers(ctx) {
		sqlexecutor.observers = nil
	}
	if sqlexecutor.isTrace {
		defer sqlexecutor.traceExit(10, observer.GetObserverId(ctx), err, time.Since(entryTime))
	}
	return err
}
