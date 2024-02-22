package main

import (
	"time"
)

type TokenBucket struct {
	capacity int
	tokens   int
	fillRate int
	lastFill time.Time
}

func (b *TokenBucket) Action() bool {
	now := time.Now()
	filled := int(now.Sub(b.lastFill).Seconds()) * b.fillRate

	b.tokens += filled
	if b.tokens > b.capacity {
		b.tokens = b.capacity
	}

	if b.tokens-1 < 0 {
		return false
	}

	b.tokens--
	b.lastFill = now
	return true
}
