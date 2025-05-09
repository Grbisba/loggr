package loggr

import (
	"os"

	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/pkg/errors"
)

var (
	defaultEncoderConfig = zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ts",
		NameKey:        "loggr",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
)

type applier struct {
	encoder    zapcore.Encoder
	zapOptions []zap.Option
	fields     []zap.Field
	cores      []CoreConfig
}

func newApplier(opts ...Option) (*applier, error) {
	defaultJSONEncoder := zapcore.NewJSONEncoder(defaultEncoderConfig)

	a := &applier{
		encoder:    defaultJSONEncoder,
		zapOptions: nil,
		fields:     nil,
		cores: []CoreConfig{
			NewCoreConfig(defaultJSONEncoder, os.Stdout, lowPriorityLevelEnablerFunc()),
			NewCoreConfig(defaultJSONEncoder, os.Stderr, highPriorityLevelEnablerFunc()),
		},
	}

	var applyErrors error

	for _, opt := range opts {
		err := opt.apply(a)
		if err != nil {
			applyErrors = multierr.Append(applyErrors, err)
		}
	}

	return a, errors.Wrap(applyErrors, "failed to apply options")
}

func (a *applier) BuildOptions() ([]zapcore.Core, []zap.Option, []zap.Field) {
	cores := make([]zapcore.Core, len(a.cores))

	for i, c := range a.cores {
		cores[i] = c.build()
	}

	return cores, a.zapOptions, a.fields
}
