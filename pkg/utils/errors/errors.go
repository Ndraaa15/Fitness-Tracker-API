package errors

import "errors"

var (
	// ErrConfigFileNotFound is the error message for when the config file is not found
	ErrConfigFileNotFound = errors.New("CONFIG FILE NOT FOUND")
)
