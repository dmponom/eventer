package logger

import "context"

type Logger interface {
	Info(ctx context.Context, msg string, fields ...map[string]interface{})
	Trace(ctx context.Context, msg string, fields ...map[string]interface{})
	Debug(ctx context.Context, msg string, fields ...map[string]interface{})
	Error(ctx context.Context, msg string, fields ...map[string]interface{})
}
