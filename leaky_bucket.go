package main

import (
	"time"
)

type LeakyBucket struct {
	capacity int
	water    int
	leakRate int
	lastLeak time.Time
}

func (b *LeakyBucket) Action() bool {
	now := time.Now()
	leaked := int(now.Sub(b.lastLeak).Seconds()) * b.leakRate

	b.water -= leaked
	if b.water < 0 {
		b.water = 0
	}

	if b.water+1 > b.capacity {
		return false
	}

	b.water++
	b.lastLeak = now
	return true
}
