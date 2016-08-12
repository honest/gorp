package gorp

import (
	"fmt"
)

//NoFieldInTypeError is a non-fatal error, when a select query returns columns that do not exist as fields in the struct it is being mapped to
type NoFieldInTypeError struct {
	TypeName        string
	MissingColNames []string
}

// NoFieldUpdateError is a non-fatal error, when there's no column need update.
type NoFieldUpdateError struct {
}

func (err *NoFieldUpdateError) Error() string {
	return fmt.Sprintf("gorp: No fields need update")
}

func (err *NoFieldInTypeError) Error() string {
	return fmt.Sprintf("gorp: No fields %+v in type %s", err.MissingColNames, err.TypeName)
}

//NonFatalError returns true if the error is non-fatal (ie, we shouldn't immediately return)
func NonFatalError(err error) bool {
	switch err.(type) {
	case *NoFieldInTypeError, *NoFieldUpdateError:
		return true
	default:
		return false
	}
}
