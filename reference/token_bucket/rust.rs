pub struct Bucket {
    pub capacity: i64,
    pub tokens: i64,
    pub refill: i64,
}

pub fn new_bucket(capacity: i64, refill: i64) -> Option<Bucket> {
    if capacity <= 0 || refill < 0 {
        return None;
    }
    Some(Bucket { capacity, tokens: capacity, refill })
}

pub fn tick(b: &Bucket) -> Bucket {
    let t = (b.tokens + b.refill).min(b.capacity);
    Bucket { capacity: b.capacity, tokens: t, refill: b.refill }
}

pub fn try_acquire(b: &Bucket, n: i64) -> (Bucket, bool) {
    if n < 1 || b.tokens < n {
        return (Bucket { capacity: b.capacity, tokens: b.tokens, refill: b.refill }, false);
    }
    (Bucket { capacity: b.capacity, tokens: b.tokens - n, refill: b.refill }, true)
}

pub fn drain(b: &Bucket) -> Bucket {
    Bucket { capacity: b.capacity, tokens: 0, refill: b.refill }
}
