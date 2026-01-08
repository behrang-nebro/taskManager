package tmanipulation

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/behrang-nebro/taskManager/internal/models"
)

func AddTask(Tasks *models.Tasks) {

	id := len(Tasks.TaskSlice) + 1
	status := false
	time := time.Now()
	var tag string
	var description string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("what is your tag?")
	fmt.Scanln(&tag)
	fmt.Println("-------------------------------------")
	fmt.Println("what is your description?")
	scanner.Scan()
	description = scanner.Text()
	fmt.Println("-------------------------------------")

	task := models.Task{
		Id:          id,
		Tag:         tag,
		Description: description,
		Time:        time,
		Status:      status,
	}

	Tasks.TaskSlice = append(Tasks.TaskSlice, task)

}
