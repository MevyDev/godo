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

func writeTasks(todoPath string, todoList taskList) (error) {
	formattedList, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(todoPath, formattedList, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var todoPath string = "todo.json"
	var todoList taskList
	todoList, err := loadTasks(todoPath)
	if err != nil {
		fmt.Println("Load error:", err)
	}
	newTask := task{Text: "main menu", Status: "todo", Difficulty: "hard"}
	todoList.Tasks = append(todoList.Tasks, newTask)
	err = writeTasks(todoPath, todoList)
	if err != nil {
		fmt.Println("Write error:", err)
	}
}
