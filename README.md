# logcheck

`logcheck` is a custom Go linter for validating log messages.

It analyzes log calls and checks them against a set of rules for message
style and safety.

Supported loggers:

-   `log/slog`
-   `go.uber.org/zap`

The linter is implemented using `golang.org/x/tools/go/analysis` and can
be used both as a standalone analyzer and as a `golangci-lint` module
plugin.

------------------------------------------------------------------------

## Implemented Rules

### 1. Log messages must start with a lowercase letter

Incorrect:

``` go
slog.Info("Starting server")
logger.Error("Failed to connect")
```

Correct:

``` go
slog.Info("starting server")
logger.Error("failed to connect")
```

### 2. Log messages must contain only English text

Incorrect:

``` go
slog.Info("запуск сервера")
logger.Error("ошибка подключения")
```

Correct:

``` go
slog.Info("starting server")
logger.Error("failed to connect")
```

### 3. Log messages must not contain special symbols or emoji

Incorrect:

``` go
slog.Info("server started!!!")
logger.Warn("something went wrong...")
logger.Debug("server started 🚀")
```

Correct:

``` go
slog.Info("server started")
logger.Warn("something went wrong")
logger.Debug("server started")
```

### 4. Log messages must not contain potentially sensitive data

Incorrect:

``` go
slog.Info("password: " + password)
logger.Debug("token=" + token)
logger.Warn("api_key=" + apiKey)
```

Correct:

``` go
slog.Info("user authenticated successfully")
logger.Debug("request completed")
logger.Info("token validated")
```

------------------------------------------------------------------------

## Features

-   AST-based detection of log calls
-   `go/types`-based detection of `zap.Logger`
-   regex-based sensitive data detection
-   configurable custom sensitive patterns via config file
-   suggested fixes for lowercase message violations
-   support for `golangci-lint` module plugin system
-   CI workflow for automated build and test

------------------------------------------------------------------------

## Project Structure

    cmd/logcheck        standalone entrypoint
    analyzer            main analyzer logic
    internal/logs       AST helpers and suggested fixes
    rules               rule implementations
    plugin              golangci-lint module plugin wrapper
    example             example target project

------------------------------------------------------------------------

## Configuration

`logcheck` supports an optional config file named `.logcheck.yml`.

Example:

``` yaml
extra_sensitive_patterns:
  - (?i)\bsession_id\b
  - (?i)\bprivate_key\b
  - (?i)\brefresh_token\b
```

These regex patterns are added to the built‑in sensitive data checks.

You can also override the config path:

``` bash
logcheck -config custom-logcheck.yml ./...
```

------------------------------------------------------------------------

## Build

Build the standalone linter:

``` bash
go build ./cmd/logcheck
```

------------------------------------------------------------------------

## Run as Standalone Tool

Windows:

``` bash
.\logcheck.exe ./example
```

Unix-like systems:

``` bash
./logcheck ./example
```

Run tests:

``` bash
go test ./...
```

------------------------------------------------------------------------

## golangci-lint Integration

This project supports `golangci-lint` through the module plugin system.

### Build custom golangci-lint

``` bash
golangci-lint custom
```

### Run the custom binary

Windows:

``` bash
.\custom-gcl.exe run ./...
```

Unix-like systems:

``` bash
./custom-gcl run ./...
```

------------------------------------------------------------------------

## Suggested Fixes

The linter provides an automatic fix for messages that start with an
uppercase letter.

Example:

``` go
slog.Info("Starting server")
```

Suggested fix:

``` go
slog.Info("starting server")
```

------------------------------------------------------------------------

## Example

Example file: `example/main.go`

You can run the linter on it and see diagnostics for invalid log
messages.

------------------------------------------------------------------------

## Design Decisions

-   `go/analysis` is used as the core analyzer framework
-   AST traversal is used to detect log calls and inspect message
    expressions
-   `go/types` is used to identify `*zap.Logger` calls
-   sensitive data detection is heuristic‑based and intentionally
    conservative
-   config support is focused on extending sensitive‑pattern detection
    without overcomplicating the analyzer

------------------------------------------------------------------------