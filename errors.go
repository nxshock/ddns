package main

import (
	"errors"
	"fmt"
)

var (
	errVarNotFound = errors.New("variable not found")
)

func wrapError(err error, a interface{}) error {
	return fmt.Errorf("%v: %v", err, a)
}
