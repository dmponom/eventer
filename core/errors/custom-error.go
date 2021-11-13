package errors

import (
	"context"
	"fmt"
)

type CustomError struct {
	Ctx        context.Context
	Err        error
	StatusCode int
	Msg        string
}

func Make(ctx context.Context, err error, statusCode int, msg string) *CustomError {
	return &CustomError{
		Ctx:        ctx,
		Err:        err,
		StatusCode: statusCode,
		Msg:        msg,
	}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("{\"error\": \"%+v\", \"msg\": \"%s\"}", e.Err, e.Msg)
}
