package log

import (
	"context"

	"github.com/vizitiuRoman/libs/logger"
	"github.com/vizitiuRoman/libs/logger/example/internal/common/adapter/meta"
	"go.uber.org/zap"
)

func ContextEncoder(ctx context.Context, log logger.Logger) logger.Logger {
	if traceID, found := ctx.Value(meta.TraceID).(string); found {
		log = log.With(zap.String(meta.TraceID, traceID))
	}

	if spanID, found := ctx.Value(meta.SpanID).(string); found {
		log = log.With(zap.String(meta.SpanID, spanID))
	}

	if component, found := ctx.Value(meta.SessionID).(string); found {
		log = log.With(zap.String(meta.SessionID, component))
	}

	return log
}
