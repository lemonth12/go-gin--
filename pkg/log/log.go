package log

import "go.uber.org/zap"

var sugar *zap.SugaredLogger

func init() {
	cfg := zap.NewProductionConfig()
	cfg.Sampling = nil
	logger, err := cfg.Build()
	logger.WithOptions()
	if err != nil {
		panic(err)
	}
	sugar = logger.Sugar()
}

func Info(args ...interface{}) {
	sugar.Info(args)
}

func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

func Error(args ...interface{}) {
	sugar.Error(args)
}

func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	sugar.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	sugar.Fatalf(template, args...)
}

func Sync() {
	sugar.Sync()
}
