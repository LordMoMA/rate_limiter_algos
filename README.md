## How to run the application

```bash
./rate_limiter
```

## Leaky Bucket and Token Bucket

The Leaky Bucket and Token Bucket are two algorithms used for rate limiting. They have different behaviors and use cases:

### Leaky Bucket: 

The Leaky Bucket algorithm is used to control rate in a fixed manner. The metaphor is of a bucket where water leaks out at a constant rate. If water (representing network packets) pours into the bucket at a rate that exceeds the leak rate, it overflows and the excess packets are discarded. If the incoming rate is less than the leak rate, the leak rate effectively becomes the incoming rate. This algorithm smooths out bursty traffic into a constant stream, but it can result in dropped packets.

### Token Bucket: 

The Token Bucket algorithm allows for bursty traffic, up to a specified maximum burst size. The metaphor is of a bucket where tokens (representing permission to send a network packet) are added at a certain rate. If the bucket is full, incoming tokens are discarded. When a process needs to send a packet, it must first take a token from the bucket. If no tokens are available, the packet must wait. This algorithm allows for bursts of traffic until the tokens are used up, after which the rate of outgoing traffic is limited to the rate of incoming tokens.

In summary, the Leaky Bucket algorithm is more suited for situations where a steady data flow is important, while the Token Bucket algorithm is more suited for situations where occasional bursts of traffic are acceptable.

## Log Results:

```bash
LeakyBucket: true
LeakyBucket: true
LeakyBucket: true
LeakyBucket: true
LeakyBucket: true
LeakyBucket: true
LeakyBucket: true
LeakyBucket: true
LeakyBucket: true
LeakyBucket: true
LeakyBucket: false
LeakyBucket: true
LeakyBucket: false
LeakyBucket: true
LeakyBucket: false
LeakyBucket: true
LeakyBucket: false
LeakyBucket: true
LeakyBucket: false
LeakyBucket: true


TokenBucket: false
TokenBucket: false
TokenBucket: true
TokenBucket: false
TokenBucket: true
TokenBucket: false
TokenBucket: true
TokenBucket: false
TokenBucket: true
TokenBucket: false
TokenBucket: true
TokenBucket: false
TokenBucket: true
TokenBucket: false
TokenBucket: true
TokenBucket: false
TokenBucket: true
TokenBucket: false
TokenBucket: true
TokenBucket: false
```

## Leaky Bucket Result Analysis

The Leaky Bucket algorithm is used to control the rate at which requests are processed. It's called the "Leaky Bucket" because it works like a physical bucket with a small leak.

In this algorithm, each request adds a certain amount of "water" to the bucket. If the bucket is full, additional requests cause the bucket to "overflow", and these requests are rejected. Over time, "water" leaks out of the bucket at a constant rate, making room for more requests.

In your case, it seems like the bucket has a capacity of 10, a leak rate of 1 per second, and you're making a request every 500 milliseconds.

Here's how the output you posted corresponds to the Leaky Bucket algorithm:

At the start, the bucket is empty. The first 10 requests (LeakyBucket: true) are accepted and add "water" to the bucket.

The 11th request causes the bucket to overflow (LeakyBucket: false), because it arrives before any "water" has had a chance to leak out.

After this, every other request is rejected. This is because you're making requests twice as fast as the leak rate. So for every two requests, one "leak" occurs. This means one request is accepted (the "water" from it is immediately leaked out), and one request causes the bucket to overflow.

This pattern continues for the rest of the requests. The Leaky Bucket algorithm ensures that the rate of outgoing requests is constant, even if the rate of incoming requests is bursty.

## Token Bucket Result Analysis

In the TokenBucket algorithm, tokens are added to the bucket at a certain fill rate. A request can only be processed if there is a token available in the bucket to "consume". If the bucket is empty (i.e., there are no tokens), the request is rejected.

In this case, the fillRate for the TokenBucket is set to 1 token per second, and we're making a request every 500 milliseconds (as indicated by the time.Sleep(500 * time.Millisecond) in the processBucket function).

So here's what's happening:

At the start, the bucket is empty, so the first request is rejected (TokenBucket: false).

After 500 milliseconds, we make another request. But the bucket is still empty because a full second hasn't passed yet to add a new token, so this request is also rejected (TokenBucket: false).

After another 500 milliseconds, a full second has passed, so a token is added to the bucket. The next request can consume this token and is therefore accepted (TokenBucket: true).

This pattern repeats for the rest of the requests. Every second request is rejected because it arrives before a new token can be added, and every other request is accepted because it arrives after a new token has been added.

This behavior illustrates how the TokenBucket algorithm can smooth out bursts of requests. Even though you're making requests twice as fast as the fill rate, only every other request is being accepted.

