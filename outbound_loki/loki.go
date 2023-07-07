package outbound_loki

import (
	"strconv"
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
	msg := "Checked " + r.Link + ", up: " + strconv.FormatBool(r.Up) + ", status code: " + strconv.Itoa(r.StatusCode) + ", duration: " + r.Duration.String() + ", error: " + r.Error.Error()
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
}
