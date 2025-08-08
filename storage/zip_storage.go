package storage

import (
	"archive/zip"
	"errors"
	"io"
	"os"
)

type ZipStorage struct {
	*Storage
}

func (z *ZipStorage) Save(data []byte) error {
	f, err := os.Create(z.GetFilename())
	if err != nil {
		return err
	}
	defer f.Close()

	zw := zip.NewWriter(f)
	defer zw.Close()

	w, err := zw.Create("data.txt")
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func (z *ZipStorage) Load() ([]byte, error) {
	r, err := zip.OpenReader(z.GetFilename())
	if err != nil {
		return nil, err
	}
	defer r.Close()

	if len(r.File) == 0 {
		return nil, errors.New("file is empty")
	}

	file := r.File[0]
	rc, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return io.ReadAll(rc)
}

func NewZipStorage(filename string) *ZipStorage {
	return &ZipStorage{
		&Storage{filename: filename},
	}
}
