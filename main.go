package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type task struct {
	Text		string	`json:"text"`
	Status		string	`json:"status"`
	Difficulty	string	`json:"difficulty"`
}

type taskList struct {
	Tasks []task `json:"tasks"`
}


func loadTasks(todoPath string) (taskList, error) {
	data, err := os.ReadFile(todoPath)
	if err != nil {
		return taskList{Tasks: []task{}}, nil
	}

	var todos taskList
	err = json.Unmarshal(data, &todos)
	return todos, err
}

func main() {
	var todoPath string = "todo.json"
	var todoList taskList
	todoList, err := loadTasks(todoPath)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(todoList.Tasks)
}
