package logger

import (
	"go.uber.org/zap"
)

type zapLogger struct {
	sugar *zap.SugaredLogger
}

func NewZapLogger(serviceName string) (Logger, error) {
	// TODO: добавить чтение env файла для выбора zap.NewDevelopment()/zap.NewProduction()
	baseLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	sugar := baseLogger.Named(serviceName).Sugar()

	return &zapLogger{sugar: sugar}, nil
}

func (l *zapLogger) Debug(args ...any) { l.sugar.Debug(args...) }
func (l *zapLogger) Info(args ...any)  { l.sugar.Info(args...) }
func (l *zapLogger) Warn(args ...any)  { l.sugar.Warn(args...) }
func (l *zapLogger) Error(args ...any) { l.sugar.Error(args...) }
func (l *zapLogger) Fatal(args ...any) { l.sugar.Fatal(args...) }

func (l *zapLogger) Debugf(format string, args ...any) { l.sugar.Debugf(format, args...) }
func (l *zapLogger) Infof(format string, args ...any)  { l.sugar.Infof(format, args...) }
func (l *zapLogger) Warnf(format string, args ...any)  { l.sugar.Warnf(format, args...) }
func (l *zapLogger) Errorf(format string, args ...any) { l.sugar.Errorf(format, args...) }
func (l *zapLogger) Fatalf(format string, args ...any) { l.sugar.Fatalf(format, args...) }

func (l *zapLogger) With(args ...any) Logger {
	return &zapLogger{sugar: l.sugar.With(args...)}
}

func (l *zapLogger) Sync() error {
	return l.sugar.Sync()
}
