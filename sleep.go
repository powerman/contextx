package contextx

import (
	"context"
	"time"
)

// Sleep is a context-aware helper for waiting logic instead of [time.Sleep]:
// it returns nil after duration, or the context error if ctx is cancelled first.
func Sleep(ctx context.Context, duration time.Duration) error {
	t := time.NewTimer(duration)
	defer t.Stop()

	select {
	case <-t.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
