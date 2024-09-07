package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joomcode/errorx"
	"github.com/spf13/viper"
)

type Response struct {
	// OK is only true if the request was successful.
	OK bool `json:"ok"`
	// Data contains the actual data of the response.
	Data interface{} `json:"data,omitempty"`
	// Error contains the error detail if the request was not successful.
	Error *ErrorResponse `json:"error,omitempty"`
	// Total is the total number of records available for pagination.
	Total int `json:"total,omitempty"`
}

type ErrorResponse struct {
	// Code is the error code. It is not status code
	Code int `json:"code"`
	// Message is the error message.
	Message string `json:"message,omitempty"`
	// Description is the error description.
	Description string `json:"description,omitempty"`
	// StackTrace is the stack trace of the error.
	// It is only returned for debugging
	StackTrace string `json:"stack_trace,omitempty"`
}

type FieldError struct {
	// Name is the name of the field that caused the error.
	Name string `json:"name"`
	// Description is the error description for this field.
	Description string `json:"description"`
}

func SendSuccessResponse(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(
		statusCode,
		Response{
			OK:   true,
			Data: data,
		},
	)
}

func SendSuccessResponseForList(ctx *gin.Context, statusCode int, data interface{}, total int) {
	ctx.JSON(
		statusCode,
		Response{
			OK:    true,
			Data:  data,
			Total: total,
		},
	)
}

func SendErrorResponse(ctx *gin.Context, err *ErrorResponse) {
	ctx.AbortWithStatusJSON(err.Code, Response{
		OK:    false,
		Error: err,
	})
}

func HandlerError(ctx *gin.Context, err error) {

	for _, exs := range Errors {
		if errorx.IsOfType(err, exs.ErrorType) {
			er := errorx.Cast(err)

			response := ErrorResponse{
				Code:    exs.StatusCode,
				Message: er.Message(),
			}

			if viper.GetBool("debug") {
				response.Description = fmt.Sprintf("Error: %v", er)
				response.StackTrace = fmt.Sprintf("%+v", errorx.EnsureStackTrace(err))
			}
			SendErrorResponse(ctx, &response)
			return
		}
	}

	HandlerError(ctx, ErrInternalServerError.New(err.Error()))
}
