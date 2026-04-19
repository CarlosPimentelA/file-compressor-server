package pool

import (
	"runtime"
	"testing"
	"time"
)

func TestWorkerPoolStart(t *testing.T) {
	goroutinesBefore := runtime.NumGoroutine()
	pool := NewPool(100, nil)
	time.Sleep(time.Millisecond * 100)
	pool.Start()
	goroutinesAfter := runtime.NumGoroutine()

	expectedGouroutines := goroutinesAfter - goroutinesBefore

	if expectedGouroutines != pool.MaxWorkers() {
		t.Errorf("Expected %d goroutines, got %d", pool.MaxWorkers(), expectedGouroutines)
	}
}

func TestWorkerPoolStop(t *testing.T) {
	pool := NewPool(100, nil)
	pool.Start()
	pool.Stop()

	if !pool.IsStopped() {
		t.Error("expected pool to be stopped")
	}
}

func TestWorkerPoolAddJob(t *testing.T) {
	test := []struct {
		name           string
		jobQueueFull   bool
		expectError    bool
		jobQueueStoped bool
	}{
		{name: "Add job successfully", jobQueueFull: false, expectError: false, jobQueueStoped: false},
		{name: "Add job to full queue", jobQueueFull: true, expectError: true, jobQueueStoped: false},
		{name: "Add job to stopped pool", jobQueueFull: false, expectError: true, jobQueueStoped: true},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			pool := NewPool(1, nil)

			if tt.jobQueueFull {
				// Fill the job queue
				pool.AddJob(&Job{})
			}

			if tt.jobQueueStoped {
				// Stop the pool
				pool.Stop()
			}

			// Try to add a job
			err := pool.AddJob(&Job{})

			if tt.expectError && err == nil {
				t.Error("Expected error but got nil")
			}

			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error, got %v", err)
			}

		})
	}
}
