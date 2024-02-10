package utils

import (
	"github.com/getsentry/sentry-go"
	"log"
	"time"
)

func SubmitSentryLog(section, function string, err error) {
	dsn := ""
	e := sentry.Init(sentry.ClientOptions{
		Dsn:           dsn,
		Debug:         true,
		EnableTracing: true,
	})
	if e != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)

	sentry.CaptureEvent(&sentry.Event{
		Message: err.Error(),
		Tags: map[string]string{
			"section": section,
			"func":    function,
		},
	})
}
