package logging

import (
	"log/slog"
)

type Email string

func (e Email) LogValue() slog.Value {
	return slog.StringValue("*******@******.***")
}
