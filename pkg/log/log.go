package log

import (
	"io"
)

type Logger interface {
	Printf(format string, args ...interface{})

	Println(args ...interface{})

	Fatalf(format string, args ...interface{})

	Fatalln(args ...interface{})
}

type logger struct {
	log Logger
}

type writerLogger interface {
	Writer() io.Writer
}

var (
	Debug = &logger{}
	Info  = &logger{}
	Stats = &logger{}
	Trace = &logger{}

	Parse    = &logger{}
	Read     = &logger{}
	Validate = &logger{}
	Optimize = &logger{}
	Write    = &logger{}
	CLI      = &logger{}
)

func SetDebugLogger(log Logger) { _ = "STUB: not implemented"; return }

func SetInfoLogger(log Logger) { _ = "STUB: not implemented"; return }

func SetStatsLogger(log Logger) { _ = "STUB: not implemented"; return }

func SetTraceLogger(log Logger) { _ = "STUB: not implemented"; return }

func SetParseLogger(log Logger) { _ = "STUB: not implemented"; return }

func SetReadLogger(log Logger) { _ = "STUB: not implemented"; return }

func SetValidateLogger(log Logger) { _ = "STUB: not implemented"; return }

func SetOptimizeLogger(log Logger) { _ = "STUB: not implemented"; return }

func SetWriteLogger(log Logger) { _ = "STUB: not implemented"; return }

func SetCLILogger(log Logger) { _ = "STUB: not implemented"; return }

func SetDefaultDebugLogger() { _ = "STUB: not implemented"; return }

func SetDefaultInfoLogger() { _ = "STUB: not implemented"; return }

func SetDefaultStatsLogger() { _ = "STUB: not implemented"; return }

func SetDefaultTraceLogger() { _ = "STUB: not implemented"; return }

func SetDefaultParseLogger() { _ = "STUB: not implemented"; return }

func SetDefaultReadLogger() { _ = "STUB: not implemented"; return }

func SetDefaultValidateLogger() { _ = "STUB: not implemented"; return }

func SetDefaultOptimizeLogger() { _ = "STUB: not implemented"; return }

func SetDefaultWriteLogger() { _ = "STUB: not implemented"; return }

func SetDefaultCLILogger() { _ = "STUB: not implemented"; return }

func SetDefaultLoggers() { _ = "STUB: not implemented"; return }

func DisableLoggers() { _ = "STUB: not implemented"; return }

func CLIEnabled() bool { _ = "STUB: not implemented"; return false }

func DebugEnabled() bool { _ = "STUB: not implemented"; return false }

func InfoEnabled() bool { _ = "STUB: not implemented"; return false }

func OptimizeEnabled() bool { _ = "STUB: not implemented"; return false }

func ParseEnabled() bool { _ = "STUB: not implemented"; return false }

func ReadEnabled() bool { _ = "STUB: not implemented"; return false }

func StatsEnabled() bool { _ = "STUB: not implemented"; return false }

func TraceEnabled() bool { _ = "STUB: not implemented"; return false }

func ValidateEnabled() bool { _ = "STUB: not implemented"; return false }

func WriteEnabled() bool { _ = "STUB: not implemented"; return false }

func (l *logger) Printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Println(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Print(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Fatalf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Fatalln(args ...interface{}) { _ = "STUB: not implemented"; return }
