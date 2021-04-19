// Package locker provides locker interface
//
package locker

import (
	"context"
	"io"
	"time"
)

type (
	Locker interface {
		LockContext(context.Context, time.Duration) error
		UnLock() error
		io.ReadWriteSeeker
		io.Closer
	}
)
