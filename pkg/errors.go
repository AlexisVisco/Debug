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
	return fmt.Sprintf("Error code %.4d, %s ", d.Code, d.Error)
}


var DebugNotFound = &DebugError{errors.New("debug structure not found"), 1}
var DebugExist = &DebugError{errors.New("debug structure already exist"), 2}