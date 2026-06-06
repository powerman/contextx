package contextx_test

import (
	"context"
	"testing"
	"time"

	"github.com/powerman/check"

	"github.com/powerman/contextx"
)

func TestSleep(tt *testing.T) {
	tt.Parallel()
	t := check.T(tt)

	// Returns nil once the duration elapses.
	t.Nil(contextx.Sleep(context.Background(), time.Millisecond))

	// Returns the context error when cancelled before the duration elapses.
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	t.Err(contextx.Sleep(cancelled, time.Hour), context.Canceled)

	// Reports a deadline as the context error too.
	expired, cancelExpired := context.WithTimeout(context.Background(), time.Nanosecond)
	tt.Cleanup(cancelExpired)
	t.Err(contextx.Sleep(expired, time.Hour), context.DeadlineExceeded)
}
