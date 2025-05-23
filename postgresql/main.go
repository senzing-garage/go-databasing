package postgresql

import (
	"context"
	"errors"

	"github.com/senzing-garage/go-observing/observer"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type Postgresql interface {
	GetCurrentWatermark(ctx context.Context) (string, int, error)
	RegisterObserver(ctx context.Context, observer observer.Observer) error
	SetLogLevel(ctx context.Context, logLevelName string) error
	SetObserverOrigin(ctx context.Context, origin string)
	UnregisterObserver(ctx context.Context, observer observer.Observer) error
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "SZSDK6423xxxx".
const ComponentID = 6423

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for postgresql implementation.
var IDMessages = map[int]string{
	1:    "Enter GetCurrentWatermark().",
	2:    "Exit  GetCurrentWatermark() returned (%s, %d, %v).",
	3:    "Enter RegisterObserver(%s).",
	4:    "Exit  RegisterObserver(%s) returned (%v).",
	5:    "Enter SetLogLevel(%s).",
	6:    "Exit  SetLogLevel(%s) returned (%v).",
	7:    "Enter UnregisterObserver(%s).",
	8:    "Exit  UnregisterObserver(%s) returned (%v).",
	8001: "GetCurrentWatermark",
	8002: "RegisterObserver",
	8003: "SetLogLevel",
	8004: "UnregisterObserver",
}

// Status strings for specific messages.
var IDStatuses = map[int]string{}

var errPackage = errors.New("sqlexecutor")
