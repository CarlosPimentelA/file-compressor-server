package pool

import (
	"compressor_server/internal/core"
	"compressor_server/internal/network"
	"fmt"
	"log"
	"time"
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

func (j *Job) Execute(compressor core.FileCompresor) error {
	defer j.Connection.Close()
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

	err = j.Connection.Write(compressedData)

	if err != nil {
		return fmt.Errorf("Write error: %w", err)
	}

	return nil
}
