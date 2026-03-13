package foo

import "log/slog"

func testSlog() {
	slog.Info("starting server")
	slog.Error("failed to connect")
	slog.Warn("something happened")
	slog.Debug("debug message")

	slog.Info("Starting server")    // want "log message must start with a lowercase letter"
	slog.Error("Failed to connect") // want "log message must start with a lowercase letter"

	slog.Info("запуск сервера")      // want "log message must contain only English text"
	slog.Error("ошибка подключения") // want "log message must contain only English text"
	slog.Warn("server запущен")      // want "log message must contain only English text"
}
