package log

import (
	"context"
	"fmt"
	zaplogfmt "github.com/sykesm/zap-logfmt"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"regexp"
)

type Logger struct {
	*zap.Logger
	zapcore.Core
	enc zapcore.Encoder
}

func New(format string) (*Logger, error) {
	var enc zapcore.Encoder
	regex := "(?i)(?:json|logfmt)"
	re := regexp.MustCompile(regex)
	logType := re.FindString(format)
	if logType == "" {
		return nil, fmt.Errorf("invalid log format: %s", format)
	}
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "T"
	config.CallerKey = "C"
	config.MessageKey = "M"
	config.LevelKey = "L"
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncodeCaller = zapcore.ShortCallerEncoder
	switch logType {
	case "json":
		enc = zapcore.NewJSONEncoder(config)
		break
	case "logfmt":
		enc = zaplogfmt.NewEncoder(config)
		break
	}
	return &Logger{
		Logger: nil,
		Core:   nil,
		enc:    enc,
	}, nil
}

func Must() *Logger {
	log, err := New("json")
	if err != nil {
		panic(err)
	}
	return log
}

func (l *Logger) Create() *Logger {
	l.Core = zapcore.NewCore(l.enc, os.Stdout, zap.DebugLevel)
	l.Logger = zap.New(l.Core, zap.AddCaller())
	return l
}

func (l *Logger) WithOpenTelemetryTraces(ctx context.Context) *Logger {
	traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
	spanID := trace.SpanFromContext(ctx).SpanContext().SpanID().String()
	fields := []zap.Field{
		zap.String("traceID", traceID),
		zap.String("spanID", spanID),
	}
	*l.Logger = *zap.New(l.Core, zap.AddCaller()).With(fields...)
	return l
}
