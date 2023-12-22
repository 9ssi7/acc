package acc

import (
	"time"
)

type Data[T interface{}] struct {
	Original T
	ID       string
}

type ProcessorFunc[T interface{}] func(data []T)

type Accumulator[T interface{}] interface {
	Add(T) (string, error)
	Cancel(string) error
	Start() error
}

type Adder[T interface{}] func(T) (string, error)

type Config[T interface{}] struct {
	Add       Adder[T]
	Processor ProcessorFunc[T]
	Storage   DataStorage[Data[T]]
	Interval  time.Duration
	StartTime time.Time
}
