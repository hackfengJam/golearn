package main

import (
	"fmt"
	"math"
	"time"
)

type Limit float64

// Inf is the infinite rate limit; it allows all events (even if burst is zero).
const Inf = Limit(math.MaxFloat64)

// Every converts a minimum time interval between events to a Limit.
func Every(interval time.Duration) Limit {
	if interval <= 0 {
		return Inf
	}
	return 1 / Limit(interval.Seconds())
}

func (limit Limit) durationFromTokens(tokens float64) time.Duration {
	seconds := tokens / float64(limit)
	return time.Nanosecond * time.Duration(1e9*seconds)
}

func main() {
	limit := Every(time.Second / time.Duration(100))
	fmt.Println(limit.durationFromTokens(-100))
}
