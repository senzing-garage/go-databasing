package connector

import "errors"

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "SZSDK6421xxxx".
const ComponentID = 6421

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var errPackage = errors.New("connector")
