package logging

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

type loggerKey struct{}

// WithLogger adds a logger to the context
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// FromContext retrieves the logger from context and enriches it with trace information
func FromContext(ctx context.Context) *slog.Logger {
	logger := slog.Default()

	// Try to get logger from context
	if ctxLogger, ok := ctx.Value(loggerKey{}).(*slog.Logger); ok {
		logger = ctxLogger
	}

	// Add trace context if available
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		logger = logger.With(
			"trace_id", span.SpanContext().TraceID().String(),
			"span_id", span.SpanContext().SpanID().String(),
		)
	}

	return logger
}

// Info logs an info message with context
func Info(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).InfoContext(ctx, msg, args...)
}

// Debug logs a debug message with context
func Debug(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).DebugContext(ctx, msg, args...)
}

// Warn logs a warning message with context
func Warn(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).WarnContext(ctx, msg, args...)
}

// Error logs an error message with context
func Error(ctx context.Context, msg string, err error, args ...any) {
	allArgs := append([]any{"error", err}, args...)
	FromContext(ctx).ErrorContext(ctx, msg, allArgs...)
}
