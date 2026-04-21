package network

import (
	"fmt"
	"io"
	"log"
	"net"
)

type TCPConnection struct {
	conn net.Conn
}

type Connection interface {
	Read() ([]byte, error)
	Write([]byte) error
	Close() error
}

func NewTCPConnection(conn net.Conn) *TCPConnection {
	return &TCPConnection{
		conn: conn,
	}
}

func (tcpC *TCPConnection) Read() ([]byte, error) {

	data, err := io.ReadAll(tcpC.conn)

	if err != nil {
		return nil, err
	}

	return Decode(data)
}

func (tcpC *TCPConnection) Write(data []byte) error {
	encodedData, err := Encode(data)

	if err != nil {
		log.Printf("Encode err: %v", err)
	}

	_, err = tcpC.conn.Write(encodedData)

	if err != nil {
		return fmt.Errorf("Error writing to connection: %v", err)
	}

	return nil
}

func (tcpC *TCPConnection) Close() error {
	return tcpC.conn.Close()
}
