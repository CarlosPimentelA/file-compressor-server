package pool

import (
	"compressor_server/internal/core"
	"fmt"
	"testing"
)

func TestJobConnectionAlwaysCloses(t *testing.T) {
	fake := &FakeConnection{
		readError: fmt.Errorf("Read Error"),
	}

	job := NewJob(fake)

	compressor := core.NewFileCompresor()

	_ = job.Execute(*compressor)
	if !fake.closed {
		t.Error("Connection was not closed")
	}
}

func TestJobWriteError(t *testing.T) {
	fake := &FakeConnection{
		readData:   []byte("Test data"),
		writeError: fmt.Errorf("Write error"),
	}

	job := NewJob(fake)
	compressor := core.NewFileCompresor()

	err := job.Execute(*compressor)

	if err == nil {
		t.Error("Expected write error but got nil")
	}
}

func TestJobExecute(t *testing.T) {
	tests := []struct {
		name      string // Nombre del caso (para identificar fallos)
		readData  []byte
		readErr   error // Entradas
		expectErr bool  // Resultado esperado
	}{
		{name: "Happy path", readData: []byte("Test data"), readErr: nil, expectErr: false},
		{name: "Read error", readData: nil, readErr: fmt.Errorf("Read error"), expectErr: true},
		{name: "Empty Data", readData: []byte{}, readErr: nil, expectErr: false},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			fake := &FakeConnection{
				readData:  tt.readData,
				readError: tt.readErr,
			}
			job := NewJob(fake)
			compressor := core.NewFileCompresor()
			err := job.Execute(*compressor)
			if tt.expectErr && err == nil {
				t.Errorf("Expected error but got nil")
			}

			if !tt.expectErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

		})
	}
}
