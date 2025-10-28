package stream

import (
	"errors"
	"sync"
)

var controlAlreadyStoppedError = errors.New("Control already stopped")

type Control interface {
	Control() <-chan struct{}
	Stop() error
}

type control struct {
	stopOnce sync.Once
	control  chan struct{}
}

func (self *control) Control() <-chan struct{} {
	return self.control
}

func (self *control) Stop() error {
	stopped := false

	self.stopOnce.Do(func() {
		stopped = true
		close(self.control)
	})

	if !stopped {
		return controlAlreadyStoppedError
	}

	return nil
}

func NewControl() *control {
	return &control{sync.Once{}, make(chan struct{})}
}
