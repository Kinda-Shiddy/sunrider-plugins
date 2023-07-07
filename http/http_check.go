package http_check

import (
	"net/http"
	"strings"
	"time"
)

type LinkCheckResult struct {
	Link     string
	Up       bool
	Duration time.Duration
	Error    error
}

// Modify the CheckLink function to use the blackbox_exporter for link checking
func CheckLink(link string, c chan LinkCheckResult, client *http.Client) {
	trimmedLink := strings.TrimSpace(link)

	// Use the blackbox_exporter URL for link checking
	blackboxURL := "http://blackbox_exporter:9115/probe?module=http_2xx&target=" + trimmedLink

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
