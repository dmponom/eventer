package logger

import (
	contextTools "eventer/core/context-tools"
	"go.uber.org/fx"

	"context"
	"eventer/config"
	"github.com/sirupsen/logrus"
	"os"
)

type logger struct {
	log *logrus.Logger

	serviceName string
}

func Make(cfg *config.Config) Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: cfg.Logger.PrettyPrint,
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.Level(cfg.Logger.DefaultLevel))

	return &logger{
		log:         log,
		serviceName: cfg.Logger.ServiceName,
	}
}

func MakeRaw(serviceName string) Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: prettyPrint,
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.Level(defaultLevel))

	return &logger{
		log:         log,
		serviceName: serviceName,
	}
}

var Module = fx.Option(
	fx.Provide(Make),
)

func (l *logger) injectFieldsToEntry(fields ...map[string]interface{}) *logrus.Entry {
	var entry *logrus.Entry
	for _, f := range fields {
		if entry == nil {
			entry = l.log.WithFields(f)
		} else {
			entry = entry.WithFields(f)
		}
	}
	return entry
}

func (l *logger) Info(ctx context.Context, msg string, fields ...map[string]interface{}) {
	requiredFields := map[string]interface{}{
		serviceNameField: l.serviceName,
		contextTraceID:   contextTools.GetTraceID(ctx),
	}
	if len(fields) == 0 {
		l.log.WithFields(requiredFields).Info(msg)
		return
	}
	l.injectFieldsToEntry(append(fields, requiredFields)...).Info(msg)
}

func (l *logger) Trace(ctx context.Context, msg string, fields ...map[string]interface{}) {
	requiredFields := map[string]interface{}{
		serviceNameField: l.serviceName,
		contextTraceID:   contextTools.GetTraceID(ctx),
	}

	if len(fields) == 0 {
		l.log.WithFields(requiredFields).Trace(msg)
		return
	}
	l.injectFieldsToEntry(append(fields, requiredFields)...).Trace(msg)
}

func (l *logger) Debug(ctx context.Context, msg string, fields ...map[string]interface{}) {
	requiredFields := map[string]interface{}{
		serviceNameField: l.serviceName,
		contextTraceID:   contextTools.GetTraceID(ctx),
	}

	if len(fields) == 0 {
		l.log.WithFields(requiredFields).Debug(msg)
		return
	}
	l.injectFieldsToEntry(append(fields, requiredFields)...).Debug(msg)
}

func (l *logger) Error(ctx context.Context, msg string, fields ...map[string]interface{}) {
	requiredFields := map[string]interface{}{
		serviceNameField: l.serviceName,
		contextTraceID:   contextTools.GetTraceID(ctx),
	}

	if len(fields) == 0 {
		l.log.WithFields(requiredFields).Error(msg)
		return
	}
	l.injectFieldsToEntry(append(fields, requiredFields)...).Error(msg)
}
