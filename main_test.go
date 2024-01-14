package basiclog

import (
	"fmt"
	"log/slog"
	"os"
	"testing"
)

func BenchmarkBasicLog(b *testing.B) {
	SetOut(os.NewFile(0, os.DevNull))
	Init()

	err := fmt.Errorf("error")

	for i := 0; i < b.N; i++ {
		Debug("", err)
		Error("", err)
	}
}

func BenchmarkSlog(b *testing.B) {
	opts := &slog.HandlerOptions{
		Level: slog.LevelError,
	}
	handler := slog.NewTextHandler(os.NewFile(0, os.DevNull), opts)

	logger := slog.New(handler)

	for i := 0; i < b.N; i++ {
		logger.Debug("error")
		logger.Error("error")
	}
}
