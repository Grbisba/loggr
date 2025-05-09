package loggr

import (
	"go.uber.org/zap/zapcore"
)

type CoreConfig struct {
	enc     zapcore.Encoder
	ws      zapcore.WriteSyncer
	enabler zapcore.LevelEnabler
}

func (c CoreConfig) build() zapcore.Core {
	return zapcore.NewCore(c.enc, c.ws, c.enabler)
}

func NewCoreConfig(enc zapcore.Encoder, ws zapcore.WriteSyncer, enabler zapcore.LevelEnabler) CoreConfig {
	return CoreConfig{enc: enc, ws: ws, enabler: enabler}
}
