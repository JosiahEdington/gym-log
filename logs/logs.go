package logs

import "context"

type Logger struct {
	Context context.Context
}

func NewLogger(ctx context.Context) *Logger {
	return &Logger{
		Context: ctx,
	}
}
