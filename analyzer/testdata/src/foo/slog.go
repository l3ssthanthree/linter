package foo

import "log/slog"

func testSlog() {
	slog.Info("starting server")    // want `found log message: "starting server"`
	slog.Error("failed to connect") // want `found log message: "failed to connect"`
	slog.Warn("something happened") // want `found log message: "something happened"`
	slog.Debug("debug message")     // want `found log message: "debug message"`
}
