# go-logging

Библиотека структурированного логирования для Go с поддержкой OpenTelemetry трейсинга.

## Возможности

- Структурированное логирование на базе `log/slog`
- Интеграция с OpenTelemetry для автоматического добавления trace_id и span_id
- Настраиваемые уровни логирования (debug, info, warn, error)
- Поддержка форматов вывода: JSON и text
- Context-based логирование

## Установка

```bash
go get github.com/kitos7/go-logging
```

## Использование

### Создание логгера

```go
import "github.com/kitos7/go-logging/logging"

// Конфигурация логгера
config := &logging.Config{
    Level:  "info",  // "debug", "info", "warn", "error"
    Format: "json",  // "json" или "text"
}

logger := logging.NewLogger(config)
```

### Context-based логирование

```go
// Добавление логгера в контекст
ctx := logging.WithLogger(context.Background(), logger)

// Использование удобных функций для логирования
logging.Info(ctx, "user logged in", "user_id", 123)
logging.Debug(ctx, "processing request", "request_id", "abc-123")
logging.Warn(ctx, "slow query detected", "duration", "5s")
logging.Error(ctx, "database connection failed", err, "host", "localhost")
```

### Интеграция с OpenTelemetry

При наличии активного span в контексте, логгер автоматически добавляет trace_id и span_id:

```go
// Если в ctx есть активный OpenTelemetry span
logging.Info(ctx, "operation completed")
// Вывод будет содержать trace_id и span_id
```

### Прямое использование логгера

```go
// Получение логгера из контекста с обогащением trace информацией
logger := logging.FromContext(ctx)
logger.Info("custom message", "key", "value")
```

## API

### Config

```go
type Config struct {
    Level  string // "debug", "info", "warn", "error"
    Format string // "json" или "text"
}
```

### Функции

- `NewLogger(config *Config) *slog.Logger` - создание нового логгера
- `WithLogger(ctx context.Context, logger *slog.Logger) context.Context` - добавление логгера в контекст
- `FromContext(ctx context.Context) *slog.Logger` - получение логгера из контекста с trace информацией
- `Info(ctx context.Context, msg string, args ...any)` - логирование info сообщения
- `Debug(ctx context.Context, msg string, args ...any)` - логирование debug сообщения
- `Warn(ctx context.Context, msg string, args ...any)` - логирование warning сообщения
- `Error(ctx context.Context, msg string, err error, args ...any)` - логирование error сообщения

## Требования

- Go 1.25.1 или выше
- go.opentelemetry.io/otel/trace v1.38.0

## Лицензия

MIT