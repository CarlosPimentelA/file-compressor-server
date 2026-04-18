package main

import (
	"compressor_server/internal/core"
	"compressor_server/internal/network"
	"compressor_server/internal/pool"
	"log"
	"net"
)

func main() {
	jobQueueSize := 100
	listener, err := net.Listen("tcp", ":4000")

	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	defer listener.Close()

	compress := core.NewFileCompresor()

	jobPool := pool.NewPool(jobQueueSize, compress)
	jobPool.Start()

	defer jobPool.Stop()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("Connection error: %v", err)
			continue
		}

		job := pool.NewJob(network.NewTCPConnection(conn))

		if err := jobPool.AddJob(job); err != nil {
			log.Printf("Add job to queue error: %v", err)
			conn.Close()
		}
	}
}
