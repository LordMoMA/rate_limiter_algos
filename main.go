package main

import (
	"fmt"
	"time"
)

type Bucket interface {
	Action() bool
}

func processBucket(b Bucket, name string) {
	for i := 0; i < 20; i++ {
		fmt.Printf("%s: %v\n", name, b.Action())
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	leakyBucket := &LeakyBucket{capacity: 10, leakRate: 1, lastLeak: time.Now()}
	processBucket(leakyBucket, "LeakyBucket")

	tokenBucket := &TokenBucket{capacity: 10, fillRate: 1, lastFill: time.Now()}
	processBucket(tokenBucket, "TokenBucket")
}
