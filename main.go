package main

import (
	"context"
	"fmt"

	"github.com/senzing/go-databasing/examplepackage"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdTemplate = "senzing-9999%04d"

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Values updated via "go install -ldflags" parameters.

var (
	programName    string = "unknown"
	buildVersion   string = "0.0.0"
	buildIteration string = "0"
)

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func exampleFunction(ctx context.Context, name string, version string, iteration string) error {
	fmt.Printf("exampleFunction: %s  %s-%s\n", programName, buildVersion, buildIteration)
	return nil
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	ctx := context.TODO()

	// Calling a function in main.go.

	exampleFunction(ctx, programName, buildVersion, buildIteration)

	// Using a package

	examplePackage := &examplepackage.ExamplePackageImpl{
		Something: " Main says 'Hi!'",
	}

	examplePackage.SaySomething(ctx)
}
