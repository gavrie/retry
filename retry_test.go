package retry_test

import (
	"testing"
	"time"

	retry "github.com/gavrie/retry"
)

func TestRetrier(t *testing.T) {
	const (
		timeout = 5 * time.Second
		retries = 3
	)

	r := retry.NewBasic(timeout, retries)
	total := r.TotalTimeout()
	expected := timeout * retries

	if total != expected {
		t.Fatalf("Unexpected total timeout."+
			" Got: %v, Expected: %v", total, expected)
	}
}

func TestExponential(t *testing.T) {
	const (
		timeout = 5 * time.Second
		retries = 3
	)

	r := retry.NewExponential(timeout, retries)
	total := r.TotalTimeout()
	expected := timeout + 2*timeout + 4*timeout

	if total != expected {
		t.Fatalf("Unexpected total timeout. Got: %v, Expected: %v", total, expected)
	}
}
