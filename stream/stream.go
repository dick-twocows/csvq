package stream

import "errors"

var SourceAlreadyStartedError = errors.New("Source already started")

type Source[T any] interface {
	Control() Control
	Out() <-chan T
	Start() error
}

type Intermediate[T, R any] interface {
	Control() Control
	Out() <-chan R
	Start() error
}

type Terminal[T any] interface {
	Control() Control
	Start() error
}
