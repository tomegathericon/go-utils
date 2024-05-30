package log4go

import (
	zaplogfmt "github.com/sykesm/zap-logfmt"
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
