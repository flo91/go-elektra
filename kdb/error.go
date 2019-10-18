package kdb

import "github.com/pkg/errors"

// error codes taken from libelektra/src/error/specification
var (
	ErrResource            = errors.New("C01100 Resource")
	ErrOutOfMemory         = errors.New("C01110 - OutOfMemory")
	ErrInstallation        = errors.New("C01200 - Installation")
	ErrInternal            = errors.New("C01310 - Internal")
	ErrInterface           = errors.New("C01320 - Interface")
	ErrPluginMisbehavior   = errors.New("C01330 - PluginMisbehavior")
	ErrConflictingState    = errors.New("C02000 - ConflictingState")
	ErrValidationSyntactic = errors.New("C03100 - ValidationSyntactic")
	ErrValidationSemantic  = errors.New("C03200 - ValidationSemantic")
)

var (
	errCodeMap = map[string]error{
		"C01110": ErrOutOfMemory,
		"C01200": ErrInstallation,
		"C01310": ErrInternal,
		"C01320": ErrInterface,
		"C01330": ErrPluginMisbehavior,
		"C02000": ErrConflictingState,
		"C03100": ErrValidationSyntactic,
		"C03200": ErrValidationSemantic,
	}
)
