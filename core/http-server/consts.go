package httpserver

import "time"

const (
	xTraceIDHeader = "x-trace-id"

	readTimeout       = 300 * time.Second
	readHeaderTimeout = 300 * time.Second
	writeTimeout      = 300 * time.Second
	idleTimeout       = 300 * time.Second
)
