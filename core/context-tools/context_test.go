package contexttools

import (
	stringTools "eventer/core/strings-tools"

	"context"
	"fmt"
	"testing"
)

func TestPopulateValues(t *testing.T) {
	var (
		traceID     = stringTools.GenerateTraceID()
		remoteAddr  = stringTools.GetRandom(10)
		emptyRecord = fmt.Sprintf("%v", nil)
	)

	tCases := []struct {
		name               string
		traceID            string
		remoteAddr         string
		expectedTraceID    string
		expectedRemoteAddr string
	}{
		{
			name:               "should all be empty",
			traceID:            emptyRecord,
			remoteAddr:         emptyRecord,
			expectedTraceID:    emptyRecord,
			expectedRemoteAddr: emptyRecord,
		},
		{
			name:               "should traceID be empty",
			traceID:            emptyRecord,
			remoteAddr:         remoteAddr,
			expectedTraceID:    emptyRecord,
			expectedRemoteAddr: remoteAddr,
		},
		{
			name:               "should remoteAddr be empty",
			traceID:            traceID,
			remoteAddr:         emptyRecord,
			expectedTraceID:    traceID,
			expectedRemoteAddr: emptyRecord,
		},
		{
			name:               "should be okay for all fields",
			traceID:            traceID,
			remoteAddr:         remoteAddr,
			expectedTraceID:    traceID,
			expectedRemoteAddr: remoteAddr,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			ctx := context.Background()
			ctx = context.WithValue(ctx, TracingTraceID, tc.traceID)
			ctx = context.WithValue(ctx, TracingClientIP, tc.remoteAddr)

			result := context.Background()
			result = PopulateValues(ctx, result)

			actualTraceID := fmt.Sprintf("%v", result.Value(TracingTraceID))
			if actualTraceID != tc.expectedTraceID {
				tt.Errorf("expected traceID = %s actual traceID = %s", tc.expectedTraceID, actualTraceID)
			}

			actualRemoteAddr := fmt.Sprintf("%v", result.Value(TracingClientIP))
			if actualRemoteAddr != tc.expectedRemoteAddr {
				tt.Errorf("expected remoteAddr = %s actual remoteAddr = %s", tc.expectedRemoteAddr, actualRemoteAddr)
			}
		})
	}
}

func TestGetTraceID(t *testing.T) {
	var (
		traceID     = stringTools.GenerateTraceID()
		emptyRecord = ""
	)

	emptyCtx := context.Background()
	ctx := context.WithValue(emptyCtx, TracingTraceID, traceID)

	tCases := []struct {
		name            string
		entryCtx        context.Context
		expectedTraceID string
	}{
		{
			name:            "GetTraceID: should traceID be empty",
			entryCtx:        emptyCtx,
			expectedTraceID: emptyRecord,
		},
		{
			name:            "should should traceID not empty",
			entryCtx:        ctx,
			expectedTraceID: traceID,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			result := GetTraceID(tc.entryCtx)
			if tc.expectedTraceID != result {
				tt.Fatalf("expected traceID = %v but actual traceID = %v", tc.expectedTraceID, result)
			}
		})
	}
}

func TestGetClientEmail(t *testing.T) {
	var (
		clientEmail = stringTools.GetRandomEmail()
		emptyRecord = ""
	)

	emptyCtx := context.Background()
	ctx := context.WithValue(emptyCtx, TracingClientEmail, clientEmail)

	tCases := []struct {
		name            string
		entryCtx        context.Context
		expectedTraceID string
	}{
		{
			name:            "GetClientEmail: should clientEmail be empty",
			entryCtx:        emptyCtx,
			expectedTraceID: emptyRecord,
		},
		{
			name:            "GetClientEmail: should should clientEmail not empty",
			entryCtx:        ctx,
			expectedTraceID: clientEmail,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			result := GetClientEmail(tc.entryCtx)
			if tc.expectedTraceID != result {
				tt.Fatalf("expected clientEmail = %v but actual clientEmail = %v", tc.expectedTraceID, result)
			}
		})
	}
}

func TestPopulateValue(t *testing.T) {
	var (
		clientEmail = stringTools.GetRandomEmail()
		emptyRecord = fmt.Sprintf("%v", nil)
	)

	tCases := []struct {
		name                string
		clientEmail         string
		expectedClientEmail string
	}{
		{
			name:                "should clientEmail be empty",
			clientEmail:         emptyRecord,
			expectedClientEmail: emptyRecord,
		},
		{
			name:                "should clientEmail be empty",
			clientEmail:         emptyRecord,
			expectedClientEmail: emptyRecord,
		},
		{
			name:                "should be okay for all fields",
			clientEmail:         clientEmail,
			expectedClientEmail: clientEmail,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			ctx := context.Background()
			ctx = PopulateValue(ctx, TracingClientEmail, tc.clientEmail)

			if email := fmt.Sprintf("%v", ctx.Value(TracingClientEmail)); email != tc.expectedClientEmail {
				tt.Errorf("expected clientEmail = %s actual clientEmail = %s", tc.expectedClientEmail, email)
			}
		})
	}
}

func TestAppendUserContext(t *testing.T) {
	var (
		clientEmail = stringTools.GetRandomEmail()
		emptyRecord = fmt.Sprintf("%v", nil)
	)

	tCases := []struct {
		name                string
		clientEmail         string
		expectedClientEmail string
	}{
		{
			name:                "should clientEmail be empty",
			clientEmail:         emptyRecord,
			expectedClientEmail: emptyRecord,
		},
		{
			name:                "should clientEmail be empty",
			clientEmail:         emptyRecord,
			expectedClientEmail: emptyRecord,
		},
		{
			name:                "should be okay for all fields",
			clientEmail:         clientEmail,
			expectedClientEmail: clientEmail,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			ctx := context.Background()
			ctx = AppendUserContext(ctx, tc.clientEmail)

			if email := fmt.Sprintf("%v", ctx.Value(TracingClientEmail)); email != tc.expectedClientEmail {
				tt.Errorf("expected clientEmail = %s actual clientEmail = %s", tc.expectedClientEmail, email)
			}
		})
	}
}
