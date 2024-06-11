package checker

import (
	"context"

	"github.com/senzing-garage/go-observing/observer"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type Checker interface {
	IsSchemaInstalled(ctx context.Context) (bool, error)
	RecordCount(ctx context.Context) (int64, error)
	RegisterObserver(ctx context.Context, observer observer.Observer) error
	SetLogLevel(ctx context.Context, logLevelName string) error
	SetObserverOrigin(ctx context.Context, origin string)
	UnregisterObserver(ctx context.Context, observer observer.Observer) error
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "senzing-6423xxxx".
const ComponentID = 6424

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for sqlfiler implementation.
var IDMessages = map[int]string{
	1:    "Enter IsSchemaInstalled().",
	2:    "Exit  IsSchemaInstalled() returned (%i).",
	3:    "Enter RegisterObserver(%s).",
	4:    "Exit  RegisterObserver(%s) returned (%v).",
	5:    "Enter SetLogLevel(%s).",
	6:    "Exit  SetLogLevel(%s) returned (%v).",
	7:    "Enter UnregisterObserver(%s).",
	8:    "Exit  UnregisterObserver(%s) returned (%v).",
	9:    "Enter RecordCount().",
	10:   "Exit  RecordCount() returned (%i).",
	8001: "IsSchemaInstalled",
	8002: "RegisterObserver",
	8003: "SetLogLevel",
	8004: "UnregisterObserver",
	8005: "RecordCount",
}

// Status strings for specific messages.
var IDStatuses = map[int]string{}
