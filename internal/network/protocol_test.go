package network

import (
	"bytes"
	"log"
	"testing"
)

func TestDecode(t *testing.T) {
	encodedInput, errInput := Encode([]byte("Hello, World!"))

	if errInput != nil {
		t.Errorf("Input err: %v", errInput)
	}

	emptyInput, _ := Encode([]byte(""))

	tests := []struct {
		name     string
		input    []byte
		expected []byte
		expErr   bool
	}{
		{
			name:     "Valid input",
			input:    encodedInput,
			expected: []byte("Hello, World!"),
			expErr:   false,
		},
		{
			name:     "Empty data",
			input:    emptyInput,
			expected: nil,
			expErr:   true,
		},
		{
			name:     "Invalid data",
			input:    nil,
			expected: nil,
			expErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Decode(tt.input)

			if err != nil && !tt.expErr {
				t.Errorf("Expected no error, got %v", err)
			}

			if err == nil && tt.expErr {
				t.Errorf("Expected error, got nil")
			}

			if !tt.expErr {
				if !bytes.Equal(tt.expected, result) {
					t.Errorf("Expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestEncode(t *testing.T) {
	input := []byte("Hello, World!")
	emptyInput := []byte("")
	encodeInput, err := Encode(input)

	if err != nil {
		log.Fatalf("Encoding err: %v", err)
	}

	tests := []struct {
		name     string
		input    []byte
		expected []byte
		wantErr  bool
	}{
		{
			name:     "Valid input",
			input:    input,
			expected: encodeInput,
			wantErr:  false,
		},
		{
			name:     "Empty data",
			input:    emptyInput,
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "Nil input",
			input:    nil,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Encode(tt.input)

			if !tt.wantErr && err != nil {
				t.Errorf("Expected no error, got %v", err)
			}

			if tt.wantErr && err == nil {
				t.Errorf("Expected error, got nil")
			}

			if !tt.wantErr {
				if !bytes.Equal(tt.expected, result) {
					t.Errorf("Expected %v, got %v", tt.expected, result)
				}
			}

			if tt.wantErr && result != nil {
				t.Errorf("Expected error, got data: %v", result)
			}
		})
	}
}
