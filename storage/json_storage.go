package storage

import "os"

type JsonStorage struct {
	*Storage
}

func (s *JsonStorage) Save(data []byte) error {
	err := os.WriteFile(s.GetFilename(), data, 0644)
	return err
}

func (s *JsonStorage) Load() ([]byte, error) {
	data, err := os.ReadFile(s.GetFilename())
	return data, err
}

func (s *JsonStorage) GetFilename() string {
	return s.filename
}

func NewJsonStorage(filename string) *JsonStorage {
	return &JsonStorage{
		&Storage{filename: filename},
	}
}
