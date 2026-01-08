package models

import "time"

type Task struct {
	Id          int       `json:"id"`
	Tag         string    `json:"tag"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
	Status      bool      `json:"status"`
}
type Tasks struct {
	TaskSlice []Task
}
