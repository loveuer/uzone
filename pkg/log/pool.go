package log

import (
	"context"
	"github.com/google/uuid"
	"github.com/loveuer/uzone/pkg/opt"
	"go.uber.org/zap"
	"sync"
)

type UzoneLogger struct {
	ctx    context.Context
	caller string
	sugar  *zap.SugaredLogger
}

func (ul *UzoneLogger) WithContext(ctx context.Context) *UzoneLogger {
	ul.ctx = ctx
	return ul
}

func (ul *UzoneLogger) gc() {
	ul.ctx = nil
	ul.caller = ""
	_ = ul.sugar.Sync()
	UzoneLoggerPool.Put(ul)
}

var UzoneLoggerPool = &sync.Pool{
	New: func() any {
		s := logger.Sugar().WithOptions(zap.AddCallerSkip(1))
		return &UzoneLogger{sugar: s}
	},
}

func New() *UzoneLogger {
	return UzoneLoggerPool.Get().(*UzoneLogger)
}

func (ul *UzoneLogger) traceId() string {
	if ul.ctx == nil {
		return uuid.Must(uuid.NewV7()).String()
	}

	if tid, ok := ul.ctx.Value(opt.TraceKey).(string); ok && tid != "" {
		return tid
	}

	return uuid.Must(uuid.NewV7()).String()
}

func (ul *UzoneLogger) With(args ...any) *UzoneLogger {
	ul.sugar = ul.sugar.With(args...)
	return ul
}

func (ul *UzoneLogger) Debug(msg string, data ...any) {
	ul.sugar.With("trace", ul.traceId()).Debugf(msg, data...)
	ul.gc()
}

func (ul *UzoneLogger) Info(msg string, data ...any) {
	ul.sugar.With("trace", ul.traceId()).Infof(msg, data...)
	ul.gc()
}

func (ul *UzoneLogger) Warn(msg string, data ...any) {
	ul.sugar.With("trace", ul.traceId()).Warnf(msg, data...)
	ul.gc()
}

func (ul *UzoneLogger) Error(msg string, data ...any) {
	ul.sugar.With("trace", ul.traceId()).Errorf(msg, data...)
	ul.gc()
}

func (ul *UzoneLogger) Panic(msg string, data ...any) {
	ul.sugar.With("trace", ul.traceId()).Panicf(msg, data...)
	ul.gc()
}

func (ul *UzoneLogger) Fatal(msg string, data ...any) {
	ul.sugar.With("trace", ul.traceId()).Fatalf(msg, data...)
	ul.gc()
}
