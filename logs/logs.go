package logs

import (
	"context"
	"time"
)

type log struct {
	LogId           int       `json:"logid"`
	Type            string    `json:"type"`
	LogName         string    `json:"logname"`
	LogValue        string    `json:"logvalue"`
	CreatedDateTime time.Time `json:"createddatetime"`
	CreatedBy       string    `json:"createdby"`
}
type Logger struct {
	Context context.Context
}

func NewLogger(ctx context.Context) *Logger {
	return &Logger{
		Context: ctx,
	}
}
