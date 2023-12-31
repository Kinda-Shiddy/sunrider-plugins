// http/http_check.go

package http_check

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type LinkCheckResult struct {
	Link     string
	Up       bool
	Duration time.Duration
	Error    error
}

func CheckLink(link string, c chan LinkCheckResult, client *http.Client) {
	trimmedLink := strings.TrimSpace(link)

	// Use the blackbox_exporter URL for link checking
	blackboxURL := "http://blackbox_exporter:9115/probe?module=http_2xx&target=" + trimmedLink
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	level.Info(logger).Log("msg", "Performing link check", "link", trimmedLink)

	req, err := http.NewRequest("GET", blackboxURL, nil)
	if err != nil {
		c <- LinkCheckResult{Link: trimmedLink, Up: false, Duration: 0, Error: err}
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		c <- LinkCheckResult{Link: trimmedLink, Up: false, Duration: 0, Error: err}
		return
	}

	defer resp.Body.Close()

	// If we reach this point, the request was successful. Send the result and return.
	c <- LinkCheckResult{Link: trimmedLink, Up: true, Duration: 0}
}
