package network

import (
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

	return data, nil
}

func (tcpC *TCPConnection) Write(data []byte) error {
	_, err := tcpC.conn.Write(data)

	if err != nil {
		log.Printf("Write error: %v", err)
	}

	return nil
}

func (tcpC *TCPConnection) Close() error {
	return tcpC.conn.Close()
}
