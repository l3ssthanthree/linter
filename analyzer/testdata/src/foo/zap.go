package foo

import (
	"fmt"

	"go.uber.org/zap"
)

func testZap(logger *zap.Logger) {
	logger.Info("starting server")
	logger.Error("failed to connect")
	logger.Warn("something happened")
	logger.Debug("debug message")

	logger.Info("Starting server")    // want "log message must start with a lowercase letter"
	logger.Error("Failed to connect") // want "log message must start with a lowercase letter"

	logger.Info("запуск сервера")      // want "log message must contain only English text"
	logger.Error("ошибка подключения") // want "log message must contain only English text"
	logger.Warn("server запущен")      // want "log message must contain only English text"

	logger.Info("server started!")         // want "log message must not contain special symbols or emoji"
	logger.Error("connection failed???")   // want "log message must not contain special symbols or emoji"
	logger.Warn("something went wrong...") // want "log message must not contain special symbols or emoji"
	logger.Debug("server started 🚀")       // want "log message must not contain special symbols or emoji"

	password := "secret"
	token := "abc"
	apiKey := "key"

	logger.Info("password: " + password)               // want "log message may contain sensitive data"
	logger.Debug("token: " + token)                    // want "log message may contain sensitive data"
	logger.Warn("api key=" + apiKey)                   // want "log message may contain sensitive data"
	logger.Info(fmt.Sprintf("password: %s", password)) // want "log message may contain sensitive data"
}
