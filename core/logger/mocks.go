package logger

import "context"

type Mock struct{}

func (*Mock) Info(ctx context.Context, msg string, fields ...map[string]interface{}) {
}

func (*Mock) Trace(ctx context.Context, msg string, fields ...map[string]interface{}) {
}

func (*Mock) Debug(ctx context.Context, msg string, fields ...map[string]interface{}) {
}

func (*Mock) Error(ctx context.Context, msg string, fields ...map[string]interface{}) {
}
