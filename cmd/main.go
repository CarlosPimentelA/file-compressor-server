package main

import (
	"compressor_server/internal/core"
	"compressor_server/internal/network"
	"compressor_server/internal/pool"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	jobQueueSize := 100
	listener, err := net.Listen("tcp", ":4000")

	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	compress := core.NewFileCompresor()

	jobPool := pool.NewPool(jobQueueSize, compress)
	jobPool.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
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
	}()
	<-quit
	log.Println("Server shutting down...")
	listener.Close()
	jobPool.Stop()
}
