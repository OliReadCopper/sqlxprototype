package logging

import (
	"io"
	"log/slog"
	"os"
)

var (
	NopLogger = slog.New(
		slog.NewJSONHandler(io.Discard, nil),
	)

	JSONLogger = slog.New(
		slog.NewJSONHandler(os.Stdout, nil),
	)
)
