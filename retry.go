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

func (r *retrier) nextTimeout() time.Duration {
	r.retries--
	return r.timeout
}

func (r *retrier) keepTrying() bool {
	return r.retries > 0
}

func (r *retrier) clone() *retrier {
	cr := *r
	return &cr
}

func (r *retrier) TotalTimeout() (total time.Duration) {
	cr := r.clone()

	for cr.keepTrying() {
		total += cr.nextTimeout()
	}
	return total
}
