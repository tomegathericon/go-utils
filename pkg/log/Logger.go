package log

import (
	"context"
	zaplogfmt "github.com/sykesm/zap-logfmt"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger struct {
	*zap.Logger
	zapcore.Core
	enc zapcore.Encoder
}

func New(format LogFormat) (*Logger, error) {
	var enc zapcore.Encoder
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "T"
	config.CallerKey = "C"
	config.MessageKey = "M"
	config.LevelKey = "L"
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncodeCaller = zapcore.ShortCallerEncoder
	switch format {
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

func Must(format LogFormat) *Logger {
	log, err := New(format)
	if err != nil {
		panic(err)
	}
	return log
}

func (l *Logger) NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "logger", l)
}

func FromContext(ctx context.Context) *Logger {
	c, _ := ctx.Value("logger").(*Logger)
	return c
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
