package acc

import (
	"testing"
	"time"
)

func TestAccumulatorWithStartTime(t *testing.T) {
	conf := Config[int]{
		Processor: func(data []int) {
			for _, d := range data {
				t.Logf("Processed Data: %d\n", d)
			}
		},
		StartTime: time.Now().Add(2 * time.Minute),
	}

	acc := New[int](conf)
	go func() {
		if err := acc.Start(); err != nil {
			t.Errorf("Error starting accumulator: %v", err)
		}
	}()

	id, err := acc.Add(42)
	if err != nil {
		t.Errorf("Error adding data: %v", err)
	}

	if err := acc.Cancel(id); err != nil {
		t.Errorf("Error canceling data: %v", err)
	}

	time.Sleep(1 * time.Second)
}

func TestAccumulatorWithInterval(t *testing.T) {
	conf := Config[int]{
		Processor: func(data []int) {
			for _, d := range data {
				t.Logf("Processed Data: %d\n", d)
			}
		},
		Interval: 1 * time.Second,
	}

	acc := New[int](conf)
	go func() {
		if err := acc.Start(); err != nil {
			t.Errorf("Error starting accumulator: %v", err)
		}
	}()

	id, err := acc.Add(42)
	if err != nil {
		t.Errorf("Error adding data: %v", err)
	}

	if err := acc.Cancel(id); err != nil {
		t.Errorf("Error canceling data: %v", err)
	}

	time.Sleep(1 * time.Second)
}
