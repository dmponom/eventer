package errors

import (
	"context"
	"net/http"
)

func InternalServerErrorBuilder(ctx context.Context, err error, msg string) error {
	return Make(ctx, err, http.StatusInternalServerError, msg)
}

func CanNotMarshalErrorBuilder(ctx context.Context, err error) error {
	return Make(ctx, err, http.StatusInternalServerError, "can not marshal struct")
}

func CanNotUnmarshalErrorBuilder(ctx context.Context, err error) error {
	return Make(ctx, err, http.StatusInternalServerError, "can not unmarshal struct")
}

func NotFoundErrorBuilder(ctx context.Context, err error, msg string) error {
	return Make(ctx, err, http.StatusNotFound, msg)
}

func BadRequestErrorBuilder(ctx context.Context, err error, msg string) error {
	return Make(ctx, err, http.StatusBadRequest, msg)
}

func UnauthorizedErrorBuilder(ctx context.Context, err error, msg string) error {
	return Make(ctx, err, http.StatusUnauthorized, msg)
}
