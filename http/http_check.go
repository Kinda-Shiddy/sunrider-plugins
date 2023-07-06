package http_check

import (
	"fmt"
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

func CheckLink(link string, c chan LinkCheckResult, client *http.Client) {
	trimmedLink := strings.TrimSpace(link)
	maxRetries := 3

	for i := 0; i < maxRetries; i++ {
		start := time.Now()

		req, err := http.NewRequest("GET", trimmedLink, nil)
		if err != nil {
			c <- LinkCheckResult{Link: trimmedLink, Up: false, Duration: time.Since(start), Error: err}
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			// Try again with HTTP/1.1
			client = &http.Client{}
			req.Proto = "HTTP/1.1"
			resp, err = client.Do(req)
			if err != nil {
				c <- LinkCheckResult{Link: trimmedLink, Up: false, Duration: time.Since(start), Error: err}
				continue
			}
		}

		defer resp.Body.Close()

		// If we reach this point, the request was successful. Send the result and return.
		c <- LinkCheckResult{Link: trimmedLink, Up: true, Duration: time.Since(start)}
		return
	}

	// If we reach this point, all retries have failed. Send a failure result.
	c <- LinkCheckResult{Link: trimmedLink, Up: false, Duration: time.Duration(maxRetries) * time.Second, Error: fmt.Errorf("all retries failed")}
}
