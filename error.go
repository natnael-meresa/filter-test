package main

import (
	"net/http"

	"github.com/joomcode/errorx"
)

type ErrorType struct {
	StatusCode int
	ErrorType  *errorx.Type
}

var Errors = []ErrorType{
	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrInvalidUserInput,
	},
	{
		StatusCode: http.StatusForbidden,
		ErrorType:  ErrAccessError,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrInternalServerError,
	},

	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrUnableToGet,
	},

	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrUnableToCreate,
	},
	{
		StatusCode: http.StatusNotFound,
		ErrorType:  ErrResourceNotFound,
	},

	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrDataAlreadyExist,
	},
	{
		StatusCode: http.StatusNotFound,
		ErrorType:  ErrNoRecordFound,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrUnableToDelError,
	},
}

var (
	databaseError    = errorx.NewNamespace("database error").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	invalidInput     = errorx.NewNamespace("validation error").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	resourceNotFound = errorx.NewNamespace("not found").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	unauthorized     = errorx.NewNamespace("unauthorized").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	AccessDenied     = errorx.RegisterTrait("You are not authorized to perform the action")
	serverError      = errorx.NewNamespace("server error")
	Unauthenticated  = errorx.NewNamespace("user authentication failed")
)

var (
	ErrUnableToCreate      = errorx.NewType(databaseError, "unable to create")
	ErrDataAlreadyExist    = errorx.NewType(databaseError, "data already exist")
	ErrUnableToGet         = errorx.NewType(databaseError, "unable to get")
	ErrUnableToUpdate      = errorx.NewType(databaseError, "unable to update")
	ErrUnableToDelError    = errorx.NewType(databaseError, "could not delete record")
	ErrInvalidUserInput    = errorx.NewType(invalidInput, "invalid user input")
	ErrResourceNotFound    = errorx.NewType(resourceNotFound, "resource not found")
	ErrAccessError         = errorx.NewType(unauthorized, "Unauthorized", AccessDenied)
	ErrInternalServerError = errorx.NewType(serverError, "internal server error")
	ErrUnExpectedError     = errorx.NewType(serverError, "unexpected error occurred")
	ErrNoRecordFound       = errorx.NewType(resourceNotFound, "no record found")
)
