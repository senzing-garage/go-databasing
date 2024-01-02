package sqlexecutor

import (
	"bufio"
	"context"

	"github.com/senzing-garage/go-observing/observer"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type SqlExecutor interface {
	ProcessFileName(ctx context.Context, filename string) error
	ProcessScanner(ctx context.Context, scanner *bufio.Scanner) error
	RegisterObserver(ctx context.Context, observer observer.Observer) error
	SetLogLevel(ctx context.Context, logLevelName string) error
	SetObserverOrigin(ctx context.Context, origin string)
	UnregisterObserver(ctx context.Context, observer observer.Observer) error
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "senzing-6422xxxx".
const ComponentId = 6422

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for sqlfiler implementation.
var IdMessages = map[int]string{
	1:    "Enter ProcessFileName(%s).",
	2:    "Exit  ProcessFileName(%s) returned (%v).",
	3:    "Enter ProcessScanner().",
	4:    "Exit  ProcessScanner() processed %d lines having %d failures. (%v).",
	5:    "Enter RegisterObserver(%s).",
	6:    "Exit  RegisterObserver(%s) returned (%v).",
	7:    "Enter SetLogLevel(%s).",
	8:    "Exit  SetLogLevel(%s) returned (%v).",
	9:    "Enter UnregisterObserver(%s).",
	10:   "Exit  UnregisterObserver(%s) returned (%v).",
	2000: "Entry: %+v",
	2001: "Processed %d lines with no failures.",
	3001: "SQL.Exec failure number %d on line %d. Error: %v",
	3002: "Processed %d lines having %d failures.",
	8001: "ProcessFileName",
	8002: "ProcessScanner.Exec",
	8003: "ProcessScanner",
	8004: "RegisterObserver",
	8005: "SetLogLevel",
	8006: "UnregisterObserver",
}

// Status strings for specific messages.
var IdStatuses = map[int]string{}
