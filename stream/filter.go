package stream

import (
	"sync"
)

type Filter[T any] interface {
	Source[T]
}

type filter[T any] struct {
	control   Control
	in        Source[T]
	startOnce sync.Once
	predicate func(Control, T) (bool, error)
	out       chan T
}

func (self *filter[T]) Control() Control {
	return self.control
}

func (self *filter[T]) Out() <-chan T {
	return self.out
}

func (self *filter[T]) Start() error {
	started := false

	self.startOnce.Do(func() {
		go func() {
			started = true

			defer func() {
				close(self.out)
			}()

			for {
				select {
				case t, ok := <-self.in.Out():
					// fmt.Printf("filter t [%v] ok [%v]\n", t, ok)
					if !ok {
						return
					}

					b, err := self.predicate(self.control, t)
					// fmt.Printf("filter b [%v] err [%v]\n", b, err)
					if err != nil {
						return
					}
					if !b {
						continue
					}

					select {
					case self.out <- t:
					case <-self.control.Control():
						return
					}

				case <-self.control.Control():
					return
				}
			}
		}()
	})

	if !started {
		return SourceAlreadyStartedError
	}

	return nil
}

func NewFilterIntermediate[T any](control Control, in Source[T], predicate func(Control, T) (bool, error)) *filter[T] {
	return &filter[T]{control, in, sync.Once{}, predicate, make(chan T)}
}
