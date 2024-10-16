// Package acc provides an accumulator mechanism to accumulate, process, and manage data.
package acc

import (
	"context"
	"time"
)

// Data represents a generic data structure with an original value and an ID.
type Data[T interface{}] struct {
	Original T
	ID       string
}

// ProcessorFunc defines the function signature for processing data.
type ProcessorFunc[T interface{}] func(ctx context.Context, data []T)

// Accumulator defines the interface for the accumulator mechanism.
type Accumulator[T interface{}] interface {
	// Add adds a new item to the accumulator and returns its ID.
	Add(context.Context, T) (string, error)
	// Cancel removes an item from the accumulator based on the provided ID.
	Cancel(context.Context, string) error
	// Start initiates the accumulator to process accumulated data.
	Start(context.Context) error
}

// Adder defines the function signature for adding items to the accumulator.
type Adder[T interface{}] func(context.Context, T) (string, error)

// Config defines the configuration structure for the accumulator mechanism.
// It encapsulates the required parameters to initialize an accumulator.
type Config[T interface{}] struct {

	// Processor is a function to process accumulated items.
	Processor ProcessorFunc[T]

	// Storage represents the storage mechanism to save and load data items.
	Storage DataStorage[Data[T]]

	// Interval specifies the time interval at which the accumulator processes accumulated data.
	Interval time.Duration

	// StartTime sets the specific time when the accumulator should start processing.
	// If set, the accumulator waits until the specified time to begin processing.
	StartTime time.Time
}
