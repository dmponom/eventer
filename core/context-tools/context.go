package contexttools

import (
	"context"
	"fmt"

	stringTools "eventer/core/strings-tools"
)

func PopulateValues(sourceCtx, targetCtx context.Context) context.Context {
	unparsedTraceID := sourceCtx.Value(TracingTraceID)
	if unparsedTraceID != nil {
		traceID := fmt.Sprintf("%v", unparsedTraceID)
		if traceID != "" {
			targetCtx = context.WithValue(targetCtx, TracingTraceID, traceID)
		}
	}

	unparsedRemoteAddr := sourceCtx.Value(TracingClientIP)
	if unparsedRemoteAddr != nil {
		remoteAddr := fmt.Sprintf("%v", unparsedRemoteAddr)
		if remoteAddr != "" {
			targetCtx = context.WithValue(targetCtx, TracingClientIP, remoteAddr)
		}
	}

	return targetCtx
}

func PopulateValue(ctx context.Context, key, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

func AppendUserContext(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, TracingClientEmail, email)
}

func AppendTracingContext(ctx context.Context, traceID, remoteAddr string) context.Context {
	if traceID == "" {
		traceID = stringTools.GenerateTraceID()
	}
	ctx = context.WithValue(ctx, TracingTraceID, traceID)
	return context.WithValue(ctx, TracingClientIP, remoteAddr)
}

func GetTraceID(ctx context.Context) string {
	if traceID, ok := ctx.Value(TracingTraceID).(string); ok {
		return traceID
	}
	return ""
}

func GetClientEmail(ctx context.Context) string {
	if clientEmail, ok := ctx.Value(TracingClientEmail).(string); ok {
		return clientEmail
	}
	return ""
}
