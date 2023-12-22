// Package acc provides an accumulator mechanism to accumulate, process, and manage data.
package acc

import (
	"errors"
	"fmt"
	"time"

	"github.com/9ssi7/nanoid"
)

// accumulator represents the core data structure that manages the accumulation process.
type accumulator[T any] struct {
	processor ProcessorFunc[T]
	storage   DataStorage[Data[T]]
	interval  time.Duration
	startTime time.Time
}

// Add adds a new item to the accumulator and returns its ID.
func (a *accumulator[T]) Add(data T) (string, error) {
	id, err := nanoid.New()
	if err != nil {
		return "", err
	}
	if err := a.storage.Save([]Data[T]{{Original: data, ID: id}}); err != nil {
		return "", err
	}
	return id, nil
}

// Cancel removes an item from the accumulator based on the provided ID.
func (a *accumulator[T]) Cancel(id string) error {
	data, err := a.storage.Load()
	if err != nil {
		return err
	}

	var newData []Data[T]
	var found bool
	for _, d := range data {
		if d.ID != id {
			newData = append(newData, d)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New(fmt.Sprintf("data not found with given ID %s", id))
	}

	if err := a.storage.Save(newData); err != nil {
		return err
	}

	return nil
}

// Start initiates the accumulator to process accumulated data based on the set interval.
// If a start time is set, it waits until that time to begin processing.
func (a *accumulator[T]) Start() error {
	if !a.startTime.IsZero() {
		for {
			now := time.Now().UTC()
			if now.Hour() == a.startTime.Hour() && now.Minute() == a.startTime.Minute() {
				break
			}
			time.Sleep(time.Minute)
		}
	}

	ticker := time.NewTicker(a.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			data, err := a.storage.Load()
			if err != nil {
				return err
			}
			if len(data) > 0 {
				var originalData []T
				for _, d := range data {
					originalData = append(originalData, d.Original)
				}
				a.processor(originalData)
			}
		}
	}
}
