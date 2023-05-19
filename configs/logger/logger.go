package logger

import (
	"github.com/rs/zerolog"
	"io"
	"os"
)

var Log zerolog.Logger

func InitLog() {
	writer := io.Writer(os.Stdout)
	Log = zerolog.New(writer)
}

func Error(from, msg string) {
	Log.Log().Timestamp().Str("level", "error").Str("from", from).Str("message", msg).Send()
}

func Info(from, msg string) {
	Log.Log().Timestamp().Str("level", "info").Str("from", from).Str("message", msg).Send()
}

func Fatal(from, msg string) {
	Log.Log().Timestamp().Str("level", "fatal").Str("from", from).Str("message", msg).Send()
	os.Exit(1)
}

func Warn(from, msg string) {
	Log.Log().Timestamp().Str("level", "warn").Str("from", from).Str("message", msg).Send()
}

func Panic(from, msg string) {
	Log.Log().Timestamp().Str("level", "panic").Str("from", from).Str("message", msg).Send()
	panic(msg)
}
