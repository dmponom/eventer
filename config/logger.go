package config

type Logger struct {
	PrettyPrint  bool   `env:"LOGGER_PRETTY_PRINT" default:"false"`
	DefaultLevel uint32 `env:"LOGGER_DEFAULT_LEVEL" default:"6"` // [6] trace level
	ServiceName  string `env:"LOGGER_PRETTY_PRINT" default:"eventer"`
}
