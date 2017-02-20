package utils

import (
	"fmt"
	"time"
)

// Retry -
func Retry(timeout time.Duration, poll time.Duration, callback func() (done bool, err error)) (err error) {

	var done bool

	timeoutAt := time.Now().Add(timeout * time.Millisecond)
	wait := poll * time.Millisecond

	for time.Now().Before(timeoutAt) {
		if done, err = callback(); done {
			return
		}
		time.Sleep(wait)
	}
	if err == nil && time.Now().After(timeoutAt) {
		err = fmt.Errorf("Last operation timed out.")
	}
	return
}
