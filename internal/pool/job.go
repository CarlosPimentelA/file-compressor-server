package pool

import (
	"fmt"
	"log"
	"time"
	"workers_server/internal/core"
	"workers_server/internal/network"
)

type Job struct {
	StartTime  time.Time
	Connection network.Connection
}

func NewJob(conn network.Connection) *Job {
	return &Job{
		StartTime:  time.Now(),
		Connection: conn,
	}
}

func (j *Job) Execute(compressor *core.FileCompresor) error {
	data, err := j.Connection.Read()

	if err != nil {
		return fmt.Errorf("conection error: %w", err)
	}

	log.Printf("Received %d in %v", len(data), time.Since(j.StartTime))

	// Add compressor logic here
	compressedData, err := compressor.Compress(data)

	if err != nil {
		return fmt.Errorf("Compress error: %w", err)
	}
	return j.finalize(compressedData)
}

func (j *Job) finalize(fileCompresed []byte) error {
	defer j.Connection.Close()
	return j.Connection.Write(fileCompresed)
}
