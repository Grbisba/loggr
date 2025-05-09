package loggr

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNew(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		t.Run("no opts", func(t *testing.T) {
			log, err := New()
			assert.NoError(t, err)
			assert.NotNil(t, log)
			log.Info("hello")
		})
		t.Run("with zap field as string", func(t *testing.T) {
			log, err := New(WithAppendToZapFieldAsString("key", "value"))
			assert.NoError(t, err)
			assert.NotNil(t, log)
			log.Info("hello")
		})
		t.Run("with zap field", func(t *testing.T) {
			log, err := New(WithZapFields(
				zap.Field{
					Key:       "key",
					Type:      zapcore.StringType,
					Integer:   1,
					String:    "value",
					Interface: nil,
				},
			))
			assert.NoError(t, err)
			assert.NotNil(t, log)
			log.Info("hello")
		})
		t.Run("with core config", func(t *testing.T) {
			log, err := New(WithCoreConfig(NewCoreConfig(
				zapcore.NewJSONEncoder(
					zapcore.EncoderConfig{
						MessageKey: "message",
					},
				), os.Stdout, zapcore.InfoLevel)),
			)
			assert.NoError(t, err)
			assert.NotNil(t, log)
			log.Info("hello")
		})
		t.Run("with encoder config", func(t *testing.T) {
			log, err := New(WithEncoderConfig(
				zapcore.EncoderConfig{
					MessageKey: "message",
					LevelKey:   "level",
				}),
			)
			assert.NoError(t, err)
			assert.NotNil(t, log)
			log.Info("hello")
		})
		t.Run("with zap option", func(t *testing.T) {
			log, err := New(WithZapOptions(
				zap.AddStacktrace(zap.InfoLevel)),
			)
			assert.NoError(t, err)
			assert.NotNil(t, log)
			log.Info("hello")
		})
		t.Run("with append to zap fields", func(t *testing.T) {
			log, err := New(
				WithZapFields(
					zap.Field{
						Key:       "key1",
						Type:      zapcore.StringType,
						Integer:   1,
						String:    "value1",
						Interface: nil,
					},
				),
				WithAppendToZapFields(
					zap.Field{
						Key:       "key",
						Type:      zapcore.StringType,
						Integer:   1,
						String:    "value",
						Interface: nil,
					},
				))
			assert.NoError(t, err)
			assert.NotNil(t, log)
			log.Info("hello")
		})
		t.Run("with append to zap option", func(t *testing.T) {
			log, err := New(
				WithZapOptions(
					zap.AddStacktrace(zap.InfoLevel),
				),
				WithAppendToZapOptions(
					zap.Fields(
						zap.Field{
							Key:       "key",
							Type:      zapcore.StringType,
							Integer:   1,
							String:    "value",
							Interface: nil,
						},
					),
				),
			)
			assert.NoError(t, err)
			assert.NotNil(t, log)
			log.Info("hello")
		})
	})
}
