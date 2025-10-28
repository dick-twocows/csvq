package stream

import (
	"fmt"
	"sync"
)

type ForEachTerminal[T any] interface {
	Terminal[T]
}

type forEachTerminal[T any] struct {
	control   Control
	in        Source[T]
	consumer  func(Control, T) error
	startOnce sync.Once
}

func (self *forEachTerminal[T]) Control() Control {
	return self.control
}

// Blocking function which receives T until Source[T] is closed or Control is closed.
func (self *forEachTerminal[T]) Start() error {
	// fmt.Printf("for each start\n")

	started := false

	self.startOnce.Do(func() {
		started = true

		for {
			select {
			case <-self.control.Control():
				return
			case t, ok := <-self.in.Out():
				// fmt.Printf("for each received t [%v]\n", t)
				if !ok {
					return
				}
				if err := self.consumer(self.control, t); err != nil {
					return
				}
			}
		}
	})

	if !started {
		return SourceAlreadyStartedError
	}

	return nil
}

func NewForEach[T any](control Control, in Source[T], consumer func(Control, T) error) ForEachTerminal[T] {
	return &forEachTerminal[T]{control, in, consumer, sync.Once{}}
}

func NewStdOutForEachConsumer[T any](c Control, t T) error {
	fmt.Printf("c [%v] t [%v]\n", c, t)
	return nil
}

// ForEach intermediate

type ForEachIntermediate[T any] interface {
	Intermediate[T, T]
}

type forEachIntermediate[T any] struct {
	control   Control
	in        Source[T]
	consumer  func(Control, T) error
	startOnce sync.Once
	out       chan T
}

func (self *forEachIntermediate[T]) Control() Control {
	return self.control
}

func (self *forEachIntermediate[T]) Out() <-chan T {
	return self.out
}

func (self *forEachIntermediate[T]) Start() error {

	started := false

	self.startOnce.Do(func() {
		started = true

		go func() {
			defer func() {
				close(self.out)
			}()

			for {
				select {
				case t, ok := <-self.in.Out():
					if !ok {
						return
					}

					err := self.consumer(self.control, t)
					if err != nil {
						return
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

func NewForEachIntermediate[T any](control Control, in Source[T], consumer func(Control, T) error) *forEachIntermediate[T] {
	return &forEachIntermediate[T]{control, in, consumer, sync.Once{}, make(chan T)}
}
