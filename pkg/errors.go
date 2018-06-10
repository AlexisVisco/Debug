package debug

import (
	"fmt"
	"errors"
)

type DebugError struct {
	Error   error
	Code    int
}

func (d DebugError) String() string {
	return fmt.Sprintf("%.4d > %s ", d.Code, d.Error)
}

type Err *DebugError

var NotFound = &DebugError{errors.New("debug structure not found"), 1}
var Exist = &DebugError{errors.New("debug structure already exist"), 2}