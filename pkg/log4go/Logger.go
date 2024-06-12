package log4go

import (
	"context"
	zaplogfmt "github.com/sykesm/zap-logfmt"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger struct {
	config zapcore.EncoderConfig
	*zap.Logger
	zapcore.Core
}

func New() *Logger {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "T"
	config.CallerKey = "C"
	config.MessageKey = "M"
	config.LevelKey = "L"
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncodeCaller = zapcore.ShortCallerEncoder
	return &Logger{
		config: config,
		Logger: nil,
		Core:   nil,
	}
}

func (l *Logger) JSON() *Logger {
	l.Core = zapcore.NewCore(zapcore.NewJSONEncoder(l.config), os.Stdout, zap.DebugLevel)
	l.Logger = zap.New(l.Core, zap.AddCaller())
	return l

}

func (l *Logger) LogFmt() *Logger {
	l.Core = zapcore.NewCore(zaplogfmt.NewEncoder(l.config), os.Stdout, zap.DebugLevel)
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
