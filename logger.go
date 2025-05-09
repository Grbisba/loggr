package loggr

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(opts ...Option) (*zap.Logger, error) {
	a, err := newApplier(opts...)
	if err != nil {
		return nil, err
	}

	cores, applierOpts, fields := a.BuildOptions()

	core := zapcore.NewTee(cores...)
	logger := zap.New(core, applierOpts...)

	for _, opt := range fields {
		logger = logger.With(opt)
	}

	zap.ReplaceGlobals(logger)
	zap.RedirectStdLog(logger)

	return logger, nil
}
