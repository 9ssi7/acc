package acc

import (
	"context"
	"testing"
	"time"
)

func TestAccumulatorWithStartTime(t *testing.T) {
	ctx := context.Background()
	acc := New(Config[int]{
		Processor: func(ctx context.Context, data []int) {
			for _, d := range data {
				t.Logf("Processed Data: %d\n", d)
			}
		},
		StartTime: time.Now().Add(2 * time.Minute),
	})
	go func() {
		if err := acc.Start(ctx); err != nil {
			t.Errorf("Error starting accumulator: %v", err)
		}
	}()

	id, err := acc.Add(ctx, 42)
	if err != nil {
		t.Errorf("Error adding data: %v", err)
	}

	if err := acc.Cancel(ctx, id); err != nil {
		t.Errorf("Error canceling data: %v", err)
	}

	time.Sleep(1 * time.Second)
}

func TestAccumulatorWithInterval(t *testing.T) {
	ctx := context.Background()
	acc := New[int](Config[int]{
		Processor: func(ctx context.Context, data []int) {
			for _, d := range data {
				t.Logf("Processed Data: %d\n", d)
			}
		},
		Interval: 1 * time.Second,
	})
	go func() {
		if err := acc.Start(ctx); err != nil {
			t.Errorf("Error starting accumulator: %v", err)
		}
	}()

	id, err := acc.Add(ctx, 42)
	if err != nil {
		t.Errorf("Error adding data: %v", err)
	}

	if err := acc.Cancel(ctx, id); err != nil {
		t.Errorf("Error canceling data: %v", err)
	}

	time.Sleep(1 * time.Second)
}
