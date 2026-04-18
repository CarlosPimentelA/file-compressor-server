package pool

import (
	"compressor_server/internal/core"
	"fmt"
	"log"
	"runtime"
	"sync"
)

type WorkerPool struct {
	maxWorkers    int
	jobQueue      chan *Job
	quit          chan struct{}
	wg            sync.WaitGroup
	fileCompresor *core.FileCompresor
}

type IWorkerPool interface {
	Start()
	Stop()
	AddJob(job *Job) error
}

func NewPool(jobQueueSize int, fileCompresor *core.FileCompresor) *WorkerPool {
	return &WorkerPool{
		maxWorkers:    runtime.NumCPU(),
		jobQueue:      make(chan *Job, jobQueueSize),
		quit:          make(chan struct{}),
		fileCompresor: fileCompresor,
		wg:            sync.WaitGroup{},
	}
}

func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for {
		select {

		case job, ok := <-wp.jobQueue:
			if !ok {
				return
			}

			if err := job.Execute(wp.fileCompresor); err != nil {
				log.Printf("Job error: %v", err)
			}

		case <-wp.quit:
			return
		}
	}
}

// Get goroutines working on the job queue
func (wp *WorkerPool) Start() {
	// Start worker goroutines
	for i := 0; i < wp.maxWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker()
		fmt.Printf("Worker %d started\n", i+1)
	}
}

// Stop the worker pool and close all connections
func (wp *WorkerPool) Stop() {
	close(wp.quit)
	wp.wg.Wait()
}

func (wp *WorkerPool) AddJob(job *Job) error {
	select {
	case wp.jobQueue <- job:
		return nil
	case <-wp.quit:
		return fmt.Errorf("The worker pool is stopping")
	default:
		return fmt.Errorf("The job queue is full")
	}
}
