package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var sugar *zap.SugaredLogger
var DEBUG = false

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)
	log := zap.New(core)
	sugar = log.Sugar()
}

func Debug(msg string) {
	if DEBUG {
		sugar.Debug(msg)
	}
}

func Debugf(template string, args ...interface{}) {
	if DEBUG {
		sugar.Debugf(template, args)
	}
}

func Info(msg string) {
	sugar.Info(msg)
}
func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args)
}

func Warn(msg string) {
	sugar.Warn(msg)
}
func Warnf(template string, args ...interface{}) {
	sugar.Warnf(template, args)
}

func Error(msg string) {
	sugar.Error(msg)
}
func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args)
}
