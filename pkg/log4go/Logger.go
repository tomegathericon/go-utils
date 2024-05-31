package log4go

import (
	"context"
	zaplogfmt "github.com/sykesm/zap-logfmt"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func GetLogger() *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "T"
	config.CallerKey = "C"
	config.MessageKey = "M"
	config.LevelKey = "L"
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncodeCaller = zapcore.ShortCallerEncoder
	logger := zap.New(zapcore.NewCore(zaplogfmt.NewEncoder(config), os.Stdout, zap.DebugLevel), zap.AddCaller())
	return logger
}

func GetLoggerWithOpenTelemetryTraces(ctx context.Context) *zap.Logger {
	logger := GetLogger()
	fields := []zap.Field{
		zap.String("traceID", trace.SpanFromContext(ctx).SpanContext().TraceID().String()),
		zap.String("spanID", trace.SpanFromContext(ctx).SpanContext().SpanID().String()),
	}
	return logger.With(fields...)
}
