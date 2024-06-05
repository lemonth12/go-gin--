package log

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

func Init(level string) {
	cfg := zap.NewProductionConfig()
	cfg.Sampling = nil
	cfg.Level = zap.NewAtomicLevel()
	err := cfg.Level.UnmarshalText([]byte(level))
	if err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	logger.WithOptions()
	if err != nil {
		panic(err)
	}
	Logger = logger.Sugar()
}
