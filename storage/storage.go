package storage

import "os"

type Storage struct {
	filename string
}

func NewStorage(filename string) *Storage {
	return &Storage{filename}
}

func (storage *Storage) Save(data []byte) error {
	err := os.WriteFile(storage.filename, data, 0644)
	return err
}

func (storage *Storage) Load() ([]byte, error) {
	data, err := os.ReadFile(storage.filename)
	return data, err
}
