package httpserver

import (
	contextTools "eventer/core/context-tools"
	"net/http"
)

func AppendTracingContext(_ http.ResponseWriter, req *http.Request) {
	_ = req.WithContext(contextTools.AppendTracingContext(req.Context(), req.Header.Get(xTraceIDHeader), req.RemoteAddr))
}
