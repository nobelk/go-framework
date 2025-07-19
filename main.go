package main

import (
	"github.com/nobelk/go-framework/logger"
	"time"
)

func main() {
	prodLogger := logger.NewLogger("Test")
	prodLogger.Infof(
		"Test log! The time is %s",
		time.Now().Format("03:04 AM"),
	)

	prodLogger.Infow("Test log!",
		"correlationId", "92e9dd6f-3d40-4ef8-baef-1ee37bc96b53",
		"userid", 123456,
	)
}
