package util

import "os"

func MustReadFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return file
}
