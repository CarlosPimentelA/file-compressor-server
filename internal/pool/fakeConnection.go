package pool

type FakeConnection struct {
	readData   []byte
	readError  error
	written    []byte
	writeError error
	closeError error
	closed     bool
}

func (fc *FakeConnection) Read() ([]byte, error) {
	return fc.readData, fc.readError
}

func (fc *FakeConnection) Write(data []byte) error {
	fc.written = data
	return fc.writeError
}

func (fc *FakeConnection) Close() error {
	fc.closed = true
	return fc.closeError
}
