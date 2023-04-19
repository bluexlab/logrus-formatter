# nested-logrus-formatter

[![Build Status](https://travis-ci.org/antonfisher/nested-logrus-formatter.svg?branch=master)](https://travis-ci.org/antonfisher/nested-logrus-formatter)
[![Go Report Card](https://goreportcard.com/badge/github.com/antonfisher/nested-logrus-formatter)](https://goreportcard.com/report/github.com/antonfisher/nested-logrus-formatter)
[![GoDoc](https://godoc.org/github.com/antonfisher/nested-logrus-formatter?status.svg)](https://godoc.org/github.com/antonfisher/nested-logrus-formatter)

Human-readable log formatter, converts _logrus_ fields to a nested structure:

![Screenshot](https://raw.githubusercontent.com/antonfisher/nested-logrus-formatter/docs/images/demo.png)

## Configuration:

```go
type Formatter struct {
	// FieldsOrder - default: fields sorted alphabetically
	FieldsOrder []string

	// UseLocalTime - use local time
	UseLocalTime bool

	// TimestampFormat - default: RFC3339Milli "2006-01-02T15:04:05.000Z07:00"
	TimestampFormat string

	// HideKeys - show [fieldValue] instead of [fieldKey:fieldValue]
	HideKeys bool

	// NoColors - disable colors
	NoColors bool

	// NoFieldsColors - apply colors only to the level, default is level + fields
	NoFieldsColors bool

	// NoFieldsSpace - no space between fields
	NoFieldsSpace bool

	// ShowFullLevel - show a full level [WARNING] instead of [WARN]
	ShowFullLevel bool

	// NoUppercaseLevel - no upper case for level value
	NoUppercaseLevel bool

	// TrimMessages - trim whitespaces on messages
	TrimMessages bool

	// CallerFirst - print caller info first
	CallerFirst bool

	// CustomCallerFormatter - set custom formatter for caller info
	CustomCallerFormatter func(*runtime.Frame) string
}
```

## Usage

```go
import (
	formatter "github.com/bluexlab/logrus-formatter"
	"github.com/sirupsen/logrus"
)

formatter.InitLogger()

logrus.Info("just info message")
// Output: 2006-01-02T15:04:05.999Z [INFO] just info message

log.WithField("component", "rest").Warn("warn message")
// Output: 2006-01-02T15:04:05.999Z [WARNING] [rest] warn message
```

See more examples in the [tests](./tests/formatter_test.go) file.

## Development

```bash
# run tests:
make test

# run demo:
make demo
```
