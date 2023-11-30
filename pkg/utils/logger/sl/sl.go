package sl

import "log/slog"

func Err(e error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(e.Error()),
	}
}
