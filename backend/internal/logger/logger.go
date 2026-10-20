package logger

import (
	"log/slog"
	"os"
)

// New builds the application's structured logger. JSON output fits
// log ingestion on Azure App Service / Log Analytics.
func New() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
