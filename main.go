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

// Returns task index if found, else -1
func findTask(todoList taskList, taskText string) (int) {
	n := len(todoList.Tasks)
	for i:=0;i<n;i++ {
		curTaskText := todoList.Tasks[i].Text
		if curTaskText == taskText {
			return i
		}
	}
	return -1
}

func addTask(todoList taskList, taskText, taskStatus, taskDifficulty string) (taskList) {
	newTask := task{Text: taskText, Status: taskStatus, Difficulty: taskDifficulty}
	taskIdx := findTask(todoList, taskText)
	if taskIdx == -1 {
		todoList.Tasks = append(todoList.Tasks, newTask)
		return todoList
	}
	todoList.Tasks[taskIdx] = newTask
	return todoList
}

func removeTask(todoList taskList, taskText string) (taskList) {
	taskIdx := findTask(todoList, taskText)
	if taskIdx == -1 {
		return todoList
	}
	n := len(todoList.Tasks)
	var updatedList = taskList{Tasks: []task{}}
	for i:=0;i<n;i++ {
		if i == taskIdx {
			continue
		}
		updatedList.Tasks = append(updatedList.Tasks, todoList.Tasks[i])
	}
	return updatedList
}

func main() {
	var todoPath string = "todo.json"

	var todoList taskList
	todoList, err := loadTasks(todoPath)
	if err != nil {
		fmt.Println("Load error:", err)
	}

	todoList = addTask(todoList, "main menu", "todo", "hard")

	todoList = removeTask(todoList, "main menu")

	err = writeTasks(todoPath, todoList)
	if err != nil {
		fmt.Println("Write error:", err)
	}
}
