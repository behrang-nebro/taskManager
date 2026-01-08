package main

import (
	"fmt"

	"github.com/behrang-nebro/taskManager/internal/iofile"
)

func main() {
	tasks, err := iofile.LoadFromFile("usersTasks.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for _, t := range tasks.TaskSlice {
		fmt.Println(t)
	}
}
