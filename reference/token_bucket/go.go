type Bucket struct{ Capacity, Tokens, Refill int64 }

func NewBucket(capacity, refill int64) (Bucket, bool) {
	if capacity <= 0 || refill < 0 {
		return Bucket{}, false
	}
	return Bucket{capacity, capacity, refill}, true
}

func Tick(b Bucket) Bucket {
	t := b.Tokens + b.Refill
	if t > b.Capacity {
		t = b.Capacity
	}
	return Bucket{b.Capacity, t, b.Refill}
}

func TryAcquire(b Bucket, n int64) (Bucket, bool) {
	if n < 1 || b.Tokens < n {
		return b, false
	}
	return Bucket{b.Capacity, b.Tokens - n, b.Refill}, true
}

func Drain(b Bucket) Bucket {
	return Bucket{b.Capacity, 0, b.Refill}
}
