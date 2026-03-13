package main

import "log/slog"

func main() {
	slog.Info("starting server")
	slog.Error("failed to connect")
}
