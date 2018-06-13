package tests

import (
	"fmt"
	"errors"
)

type ErrorDebug struct {
	Error   error
	Code    int
}

// String format the output of a ErrorDebug
func (d ErrorDebug) String() string {
	return fmt.Sprintf("Err %.4d: %s ", d.Code, d.Error)
}

type Err *ErrorDebug

// NotFound is the error returned if a debug is not found in the registry.
var NotFound = &ErrorDebug{errors.New("debug structure not found"), 1}
// Exist is the error returned if a debug already exist in the registry.
var Exist = &ErrorDebug{errors.New("debug structure already exist"), 2}