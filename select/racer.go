package racer

import (
	"net/http"
	"time"
)

func Racer(first, second string) string {
	firstDuration := measureResponseTime(first)
	secondDuration := measureResponseTime(second)

	if firstDuration < secondDuration {
		return first
	}
	return second
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
