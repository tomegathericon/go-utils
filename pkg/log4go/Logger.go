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
	*zap.Logger
}

func NewLogger() *Logger {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "T"
	config.CallerKey = "C"
	config.MessageKey = "M"
	config.LevelKey = "L"
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncodeCaller = zapcore.ShortCallerEncoder
	logger := zap.New(zapcore.NewCore(zaplogfmt.NewEncoder(config), os.Stdout, zap.DebugLevel), zap.AddCaller())
	return &Logger{
		Logger: logger,
	}
}

func (l *Logger) WithOpenTelemetryTraces(ctx context.Context) *Logger {
	fields := []zap.Field{
		zap.String("traceID", trace.SpanFromContext(ctx).SpanContext().TraceID().String()),
		zap.String("spanID", trace.SpanFromContext(ctx).SpanContext().SpanID().String()),
	}
	l.Logger = l.Logger.With(fields...)
	return l
}
