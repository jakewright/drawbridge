package retry

import (
	"math"
	"math/rand"
	"time"
)

type BackoffStrategy interface {
	// Calculate takes the number of times the request has been retried and returns the backoff
	// period in ms. If the request has only been tried once, then retryCount will be 0.
	Calculate(attempts int) time.Duration
}

// ExponentialBackoff returns backoff periods that increase exponentially 1s, 2s, 4s, ...
// Jitter is introduced by picking a random number between 0 and the backoff period.
type ExponentialBackoff struct{}

func (b *ExponentialBackoff) Calculate(attempts int) time.Duration {
	upperBound := 1000 * int(math.Pow(2, float64(attempts-1)))
	backoff := rand.Intn(upperBound)
	return time.Duration(backoff) * time.Millisecond
}
