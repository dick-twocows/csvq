package stream

import (
	"sync"
)

type supplierSource[T any] struct {
	control   Control
	startOnce sync.Once
	supplier  func(Control) (T, bool, error)
	out       chan T
}

func (self *supplierSource[T]) Control() Control {
	return self.control
}

func (self *supplierSource[T]) Out() <-chan T {
	return self.out
}

func (self *supplierSource[T]) Start() error {
	self.startOnce.Do(
		func() {
			go func() {
				defer func() {
					// fmt.Printf("supplier closing\n")
					close(self.out)
				}()
				for {
					t, ok, err := self.supplier(self.control)
					// fmt.Printf("t [%v] ok [%v] err [%v]\n", t, ok, err)
					if err != nil {
						return
					}
					if !ok {
						return
					}
					// fmt.Printf("sending t [%v]\n", t)
					select {
					case self.out <- t:
						// fmt.Printf("supplier sent t [%v]\n", t)
					case <-self.control.Control():
						return
					}
				}
			}()
		})

	return nil
}

func NewSupplierSource[T any](control Control, supplier func(Control) (T, bool, error)) Source[T] {

	return &supplierSource[T]{control, sync.Once{}, supplier, make(chan T)}
}
