package loggr

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func highPriorityLevelEnablerFunc() zap.LevelEnablerFunc {
	return func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	}
}

func lowPriorityLevelEnablerFunc() zap.LevelEnablerFunc {
	return func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	}
}

type optionFunc func(a *applier) error

func (f optionFunc) apply(a *applier) error {
	if a == nil {
		return nil
	}
	return f(a)
}

type Option interface {
	apply(*applier) error
}

func WithEncoderConfig(config zapcore.EncoderConfig) Option {
	return optionFunc(func(a *applier) error {
		a.encoder = zapcore.NewJSONEncoder(config)
		return nil
	})
}

func WithZapOptions(opts ...zap.Option) Option {
	return optionFunc(func(a *applier) error {
		a.zapOptions = opts
		return nil
	})
}

func WithAppendToZapOptions(opts ...zap.Option) Option {
	return optionFunc(func(a *applier) error {
		a.zapOptions = append(a.zapOptions, opts...)
		return nil
	})
}

func WithZapFields(fields ...zap.Field) Option {
	return optionFunc(func(a *applier) error {
		a.fields = fields
		return nil
	})
}

func WithAppendToZapFields(fields ...zap.Field) Option {
	return optionFunc(func(a *applier) error {
		a.fields = append(a.fields, fields...)
		return nil
	})
}

func WithAppendToZapFieldAsString(key string, value string) Option {
	return optionFunc(func(a *applier) error {
		a.fields = append(a.fields, zap.String(key, value))
		return nil
	})
}

func WithCoreConfig(cfg ...CoreConfig) Option {
	return optionFunc(func(a *applier) error {
		a.cores = cfg
		return nil
	})
}
