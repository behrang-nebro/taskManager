package iofile

import (
	"fmt"
	"os"
)

func CreateFile(fileName string) (*os.File, error) {
	_, err := os.Stat(fileName)

	if err == nil {
		file, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		fmt.Println("---file was loaded succesfuly---")
		return file, nil

	} else {
		fmt.Println("--creating file---")
		file, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}
		fmt.Println("---file was created---")
		return file, nil

	}
}
