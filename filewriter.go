package main

import (
	"fmt"
	"path"

	"github.com/spf13/afero"
)

type fileWriter struct {
	fs       afero.Fs
	location string
	data     []string
}

func (f fileWriter) Write() error {
	var err error
	for i, d := range f.data {
		err = f.writeFile(f.createPath(i), d)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f fileWriter) writeFile(filePath, contents string) error {
	file, err := f.fs.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(contents)
	if err != nil {
		return err
	}
	return file.Sync()
}

func (f fileWriter) createPath(i int) string {
	return path.Join(f.location, createFileName(i))
}

func createFileName(i int) string {
	return fmt.Sprintf("%03d", i)
}
