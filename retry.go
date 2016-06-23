package retry

import "time"

type retrier struct {
	timeout time.Duration
	retries int
}

func New(timeout time.Duration, retries int) *retrier {
	return &retrier{
		timeout: timeout,
		retries: retries,
	}
}

func (r *retrier) TotalTimeout() (total time.Duration) {
	for i := 0; i < r.retries; i++ {
		total += r.timeout
	}
	return total
}
