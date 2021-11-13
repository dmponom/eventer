package contexttools

type Tracing int

const (
	TracingTraceID Tracing = iota
	TracingClientIP
	TracingClientEmail
)

func (t Tracing) String() string {
	return [...]string{"traceID", "clientIP", "clientEmail"}[t]
}
