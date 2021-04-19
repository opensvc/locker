// Package locker provides locker interface
//
package locker

//go:generate mockgen -source=main.go -destination=./mock_locker/main.go

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
