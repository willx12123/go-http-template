package attr

import "log/slog"

func Err(err error) slog.Attr {
	return slog.String("err", err.Error())
}

func Uint(key string, v uint) slog.Attr {
	return slog.Uint64(key, uint64(v))
}

func Int(key string, v int) slog.Attr {
	return slog.Int64(key, int64(v))
}
