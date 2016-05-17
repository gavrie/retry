package retry

import "time"

type basicRetrier struct {
	timeout time.Duration
	retries int
}

func NewBasic(timeout time.Duration, retries int) *basicRetrier {
	return &basicRetrier{
		timeout: timeout,
		retries: retries,
	}
}

func (br *basicRetrier) nextTimeout() time.Duration {
	br.retries--
	return br.timeout
}

func (br *basicRetrier) keepTrying() bool {
	return br.retries > 0
}

func (br *basicRetrier) clone() *basicRetrier {
	cr := *br
	return &cr
}

func (br *basicRetrier) TotalTimeout() (total time.Duration) {
	cr := br.clone()

	for cr.keepTrying() {
		total += cr.nextTimeout()
	}
	return total
}

////////////////////////////////////////////////////////////

type exponentialRetrier struct {
	basicRetrier
}

func NewExponential(timeout time.Duration, retries int) *exponentialRetrier {
	er := &exponentialRetrier{}
	// Need to use less convenient syntax here due to the way embedded fields work
	er.timeout = timeout
	er.retries = retries
	return er
}

func (er *exponentialRetrier) nextTimeout() time.Duration {
	er.retries--
	t := er.timeout
	er.timeout *= 2
	return t
}

func (er *exponentialRetrier) clone() *exponentialRetrier {
	cr := *er
	return &cr
}

func (er *exponentialRetrier) TotalTimeout() (total time.Duration) {
	cr := er.clone()

	for cr.keepTrying() {
		total += cr.nextTimeout()
	}
	return total
}
