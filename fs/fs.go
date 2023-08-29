package fs

import (
	"os"
)

func CreateFile (filepath string) (string, error) {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	return file.Name(), nil
}

func CreateDirectoryIfNotExists (dirName string) {
	err := os.MkdirAll(dirName, os.ModePerm)

	if (err != nil) {
		panic(err)
	}
}