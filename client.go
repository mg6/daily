package main

import (
	"context"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/cyp0633/libcaldora/davclient"
)

// Find calendars using automatic discovery.
func GetCalendars(settings CalDAVConfig) map[string]davclient.CalendarInfo {
	calendars, err := davclient.FindCalendars(context.Background(), settings.URL, settings.Username, settings.Password)
	if err != nil {
		log.Fatal(err)
	}

	result := make(map[string]davclient.CalendarInfo)
	for _, cal := range calendars {
		result[cal.Name] = cal
	}
	return result
}

func GetDefaultLogger() *slog.Logger {
	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelInfo)

	handler := slog.NewTextHandler(io.Writer(os.Stderr), &slog.HandlerOptions{
		Level: lvl,
	})

	logger := slog.New(handler)
	return logger
}
