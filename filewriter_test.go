package main

import (
	"path"
	"testing"

	"github.com/spf13/afero"
)

func TestFileWriter(t *testing.T) {
	var (
		tmpDir   = "random"
		fs       afero.Fs
		b        []byte
		err      error
		filePath string
		fw       fileWriter
	)

	tests := [][]string{
		[]string{"hello", "world"},
	}

	for _, d := range tests {
		fs = afero.NewMemMapFs()
		fw = fileWriter{
			fs:       fs,
			location: afero.GetTempDir(fs, tmpDir),
			data:     d,
		}
		err = fw.Write()

		if err != nil {
			t.Fatalf("Error writing files\n%v", err)
		}

		for i, item := range fw.data {
			filePath = path.Join(afero.GetTempDir(fs, tmpDir), createFileName(i))
			b, err = afero.ReadFile(fs, filePath)
			if err != nil {
				t.Fatalf("Error writing file\n%v", err)
			}
			if string(b) != item {
				t.Errorf("repeatText\nExpected: %q\nGot: %q", item, string(b))
			}
		}
	}
}

func TestCreateFileName(t *testing.T) {
	var res string
	tests := []struct {
		input    int
		expected string
	}{
		{1, "001"},
		{45, "045"},
	}

	for _, fixt := range tests {
		res = createFileName(fixt.input)
		if res != fixt.expected {
			t.Errorf("createFileName\nExpected:%q\n, Got:%q", fixt.expected, res)
		}
	}
}
