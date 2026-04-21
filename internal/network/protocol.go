package network

import (
	"encoding/binary"
	"fmt"
)

func Encode(data []byte) ([]byte, error) {

	if data == nil {
		return nil, fmt.Errorf("data can't be nil!")
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("data can't be empty!")
	}
	// Get data size
	dataLen := len(data)

	header := make([]byte, 4)

	// Get total size (4 bytes for length + data)
	binary.BigEndian.PutUint32(header, uint32(dataLen))

	return append(header, data...), nil
}

func Decode(data []byte) ([]byte, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("data too short")
	}

	dataLen := binary.BigEndian.Uint32(data[:4])

	if int(dataLen) != len(data)-4 {
		return nil, fmt.Errorf("Data length mismatch: expected %d, got %d", dataLen, len(data)-4)
	}

	return data[4:], nil
}
