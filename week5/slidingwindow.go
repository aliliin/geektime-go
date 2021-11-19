package sliding

import (
	"sync"
	"time"
)

type Number struct {
	mu sync.RWMutex

	winTime     time.Duration
	bucketTime  time.Duration
	bucketCount int
	buckets     []*numberBucket
}

type numberBucket struct {
	time  int64
	value float64
}

func NewNumber(winTime time.Duration, bucketTime time.Duration) *Number {
	n := new(Number)
	n.winTime = winTime
	n.bucketTime = bucketTime
	n.bucketCount = int(n.winTime.Milliseconds() / n.bucketTime.Milliseconds())
	n.buckets = make([]*numberBucket, 0)

	return n
}

func (n *Number) Increment(i float64) {
	if i == 0 {
		return
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	b := n.getCurrentBucket()
	b.value += i
	n.removeOldBuckets()
}

func (n *Number) UpdateMax(i float64) {
	n.mu.Lock()
	defer n.mu.Unlock()

	b := n.getCurrentBucket()
	if i > b.value {
		b.value = i
	}
	n.removeOldBuckets()
}

func (n *Number) Sum(now time.Time) float64 {
	sum := float64(0)

	n.mu.RLock()
	defer n.mu.RUnlock()

	for _, b := range n.buckets {
		if b.time >= now.UnixNano()/1e6-int64(n.winTime) {
			sum += b.value
		}
	}

	return sum
}

func (n *Number) Avg(now time.Time) float64 {
	return n.Sum(now) / float64(n.bucketCount)
}

func (n *Number) Max(now time.Time) float64 {
	var max float64

	n.mu.RLock()
	defer n.mu.RUnlock()

	for _, b := range n.buckets {
		if b.time >= now.UnixNano()/1e6-int64(n.winTime) {
			if b.value > max {
				max = b.value
			}
		}
	}

	return max
}

func (n *Number) getCurrentBucket() *numberBucket {
	now := time.Now().UnixNano() / 1e6

	var b *numberBucket
	if len(n.buckets) == 0 {
		b = new(numberBucket)
		b.time = now
		n.buckets = append(n.buckets, b)
	}

	b = n.buckets[len(n.buckets)-1]
	if now-b.time >= int64(n.bucketTime) {
		b = new(numberBucket)
		b.time = now
		n.buckets = append(n.buckets, b)
	}

	return b
}

func (n *Number) removeOldBuckets() {
	now := time.Now().UnixNano()/1e6 - int64(n.winTime)

	for _, b := range n.buckets {
		if b.time <= now {
			n.buckets = n.buckets[1:]
		}
	}
}
