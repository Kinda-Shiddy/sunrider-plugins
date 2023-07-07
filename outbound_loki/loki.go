package outbound_loki

import (
	"fmt"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type LinkCheckResult struct {
	Link       string
	Up         bool
	StatusCode int
	Duration   time.Duration
	Error      error
}

type LokiLogger struct {
	logger log.Logger
}

func NewLokiLogger(logger log.Logger) *LokiLogger {
	return &LokiLogger{
		logger: logger,
	}
}

func (ll *LokiLogger) LogResult(r LinkCheckResult) {
	msg := fmt.Sprintf("Checked %s, up: %t, status code: %d, duration: %v, error: %v", r.Link, r.Up, r.StatusCode, r.Duration, r.Error)
	level.Info(ll.logger).Log(
		"msg", msg,
		"link", r.Link,
		"up", r.Up,
		"status_code", r.StatusCode,
		"duration", r.Duration,
	)
	if r.Error != nil {
		level.Error(ll.logger).Log("msg", "Error during check", "err", r.Error)
	}
	fmt.Println(msg) // Print each test result on a new line
}
