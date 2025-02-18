package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)

	WithComponent(context.Context, string) Logger
	WithMethod(context.Context, string) Logger

	AddContext(ctx context.Context) Logger
	With(fields ...zap.Field) Logger
	UnZap() *zap.Logger
	Zap(*zap.Logger)
}

type logger struct {
	*zap.Logger
	contextEncoder ContextEncoder
}

func NewLogger(cfg *Config, opts ...Option) (Logger, error) {
	zapConfig := defaultJSONConfig(cfg.Level)

	zapConfig.Encoding = cfg.Encoding

	var le zapcore.LevelEncoder
	_ = le.UnmarshalText([]byte(cfg.LevelEncoder))
	zapConfig.EncoderConfig.EncodeLevel = le

	baseLogger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}
	defer func() { _ = baseLogger.Sync() }()

	if cfg.CallerSkip != nil {
		baseLogger.WithOptions(zap.AddCallerSkip(*cfg.CallerSkip))
	}

	l := &logger{Logger: baseLogger}

	for _, opt := range opts {
		opt(l)
	}

	return l.With(zap.Int("casino.id", cfg.CasinoID)), nil
}

func NewDebugConfig() *Config {
	return &Config{
		CasinoID:     0,
		Level:        "debug",
		Encoding:     "console",
		LevelEncoder: "color",
		CallerSkip:   nil,
	}
}

func defaultJSONConfig(lvl string) *zap.Config {
	return &zap.Config{
		Level:            unmarshalLevel(lvl),
		DisableCaller:    false,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		Encoding:         "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "ts_date",
			NameKey:        "logger",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		},
	}
}

func unmarshalLevel(lvl string) zap.AtomicLevel {
	level := zap.NewAtomicLevel()
	err := level.UnmarshalText([]byte(lvl))
	if err != nil || lvl == "" {
		level.SetLevel(zap.DebugLevel)
	}

	return level
}

func (l *logger) AddContext(ctx context.Context) Logger {
	if l.contextEncoder != nil {
		return &logger{
			Logger:         l.contextEncoder(ctx, l).UnZap(),
			contextEncoder: l.contextEncoder,
		}
	}

	return l
}

func (l *logger) WithComponent(ctx context.Context, componentName string) Logger {
	const componentKey = "go.component"

	return &logger{
		Logger:         l.AddContext(ctx).With(zap.String(componentKey, componentName)).UnZap(),
		contextEncoder: l.contextEncoder,
	}
}

func (l *logger) WithMethod(ctx context.Context, methodName string) Logger {
	const methodNameKey = "go.method"

	return &logger{
		Logger:         l.AddContext(ctx).With(zap.String(methodNameKey, methodName)).UnZap(),
		contextEncoder: l.contextEncoder,
	}
}

func (l *logger) With(fields ...zap.Field) Logger {
	return &logger{
		Logger:         l.Logger.With(fields...),
		contextEncoder: l.contextEncoder,
	}
}

func (l *logger) UnZap() *zap.Logger {
	return l.Logger
}

func (l *logger) Zap(z *zap.Logger) {
	l.Logger = z
}
