// Package acc provides an accumulator mechanism to accumulate, process, and manage data.
package acc

import "time"

// New initializes a new accumulator instance based on the provided configuration.
// If no storage is provided, it defaults to an in-memory storage.
// If no interval is provided, it defaults to 15 minutes.
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
