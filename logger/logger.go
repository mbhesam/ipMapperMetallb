package logger

import (
	"log/slog"
	"os"
)

func GetLogger() *slog.Logger {
	var jhandler *slog.JSONHandler = slog.NewJSONHandler(os.Stderr, nil)
	var webjsonlog *slog.Logger = slog.New(jhandler)
	return webjsonlog
}
