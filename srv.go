package acc

import "time"

func New[T any](cnf Config[T]) Accumulator[T] {
	if cnf.Storage == nil {
		cnf.Storage = NewInMemoryStorage[T]()
	}

	if cnf.Interval == 0 {
		cnf.Interval = 15 * time.Minute
	}
	return &accumulator[T]{
		add:       cnf.Add,
		processor: cnf.Processor,
		storage:   cnf.Storage,
		interval:  cnf.Interval,
		startTime: cnf.StartTime,
	}
}
