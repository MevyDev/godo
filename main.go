package main

import (
	"fmt"
	"encoding/json"
	"os"
	"sort"
)

type task struct {
	Text		string	`json:"text"`
	Status		string	`json:"status"`
	Difficulty	int	`json:"difficulty"`
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

func addTask(todoList taskList, taskText, taskStatus string, taskDifficulty int) (taskList) {
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

func sortTasks(todoList taskList, sortOn string, sortDescending bool) (taskList, error) {
	switch sortOn {
	case "text":
		sort.Slice(todoList.Tasks, func(i, j int) bool {
		    return todoList.Tasks[i].Text < todoList.Tasks[j].Text
		})
	case "difficulty":
		sort.Slice(todoList.Tasks, func(i, j int) bool {
			return todoList.Tasks[i].Difficulty < todoList.Tasks[j].Difficulty
		})
	default:
		return todoList, fmt.Errorf("invalid sort key: %s", sortOn)
	}

	if !sortDescending {
		return todoList, nil
	}

	
	n := len(todoList.Tasks)
	var reversedTodoList = taskList{Tasks: []task{}}
	for i:=n-1;i>=0;i-- {
		reversedTodoList.Tasks = append(reversedTodoList.Tasks, todoList.Tasks[i])
	}
	return reversedTodoList, nil
}

func main() {
	var todoPath string = "todo.json"

	var todoList taskList
	todoList, err := loadTasks(todoPath)
	if err != nil {
		fmt.Println("Load error:", err)
	}

	todoList = addTask(todoList, "main menu", "todo", 5)

	todoList, err = sortTasks(todoList, "text", false)
	if err != nil {
		fmt.Println("Load error:", err)
	}

	err = writeTasks(todoPath, todoList)
	if err != nil {
		fmt.Println("Write error:", err)
	}
}
