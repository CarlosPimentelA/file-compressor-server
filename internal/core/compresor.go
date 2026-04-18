package core

type IFileCompresor interface {
	Compress(data []byte) ([]byte, error)
	Decompress(data []byte) ([]byte, error)
}

type FileCompresor struct{}

func NewFileCompresor() *FileCompresor {
	return &FileCompresor{}
}

func (fc *FileCompresor) Compress(data []byte) ([]byte, error) {
	return nil, nil
}

func (fc *FileCompresor) Decompress(data []byte) ([]byte, error) {
	return nil, nil
}
