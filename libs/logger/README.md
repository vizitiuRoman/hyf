# Logger

Logger is a logging library based on Go's `zap` package. It provides a simple and efficient way to log messages in your Go applications.

## Features

- Different log levels: Debug, Info, Warn, Error.
- Contextual logging: Add context to your logs.
- Component and method logging: Log with specific component or method names.
- Customizable: Configure the logger according to your needs.

## Usage

First, create a new logger instance:

```go
config := logger.Config{
	Level:        "debug",
	Encoding:     "json",
	LevelEncoder: "console",
}
log, err := logger.NewLogger(&config)
if err != nil {
	log.Fatal("failed to create logger")
}
```

Then, you can log messages at different levels:

```go
log.Debug("debug message")
log.Info("info message")
log.Warn("warning message")
log.Error("error message")
```

## Installation

To use Logger in your Go project, add it to your `go.mod` file:

```go
go get github.com/vizitiuRoman/libs/logger
```
