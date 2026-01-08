package iofile

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/behrang-nebro/taskManager/internal/models"
)

func CreateFile(fileName string) (*os.File, error) {
	_, err := os.Stat(fileName)

	if err == nil {
		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return nil, err
		}
		fmt.Println("---file was loaded succesfuly---")
		return file, nil

	} else {
		fmt.Println("---creating file---")
		file, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}
		fmt.Println("---file was created---")
		return file, nil

	}
}
func SaveToFile(file *os.File, tasks *models.Tasks) {
	file.Truncate(0)
	file.Seek(0, 0)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err := encoder.Encode(tasks)
	if err != nil {
		panic(err)
	}
}
func LoadFromFile(fileName string) (models.Tasks, error) {

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)

	var tasks models.Tasks

	if err != nil {
		return tasks, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(&tasks)

	return tasks, nil
}
