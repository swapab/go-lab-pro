package main

import (
	"encoding/json"
	"os"
	"strconv"
)

type Task struct {
	Title string
	Done  bool
}

func load() []Task {
	file, _ := os.ReadFile("tasks.json")
	var tasks []Task
	_ = json.Unmarshal(file, &tasks)
	return tasks
}

func save(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	_ = os.WriteFile("tasks.json", data, 0644)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		println("Usage: todo-cli <add|list|done> [task title]")
		return
	}

	command := args[1]
	tasks := load()

	switch command {
	case "add":
		if len(args) < 3 {
			println("Please provide a task title.")
			return
		}
		title := args[2]
		tasks = append(tasks, Task{Title: title})
		save(tasks)
		println("Task added:", title)

	case "list":
		for i, task := range tasks {
			status := " "
			if task.Done {
				status = "x"
			}
			println(i+1, "[", status, "]", task.Title)
		}

	case "done":
		if len(args) < 3 {
			println("Please provide the task number to mark as done.")
			return
		}
		idx, _ := strconv.Atoi(args[2])
		if idx >= 0 && idx < len(tasks) {
			tasks[idx].Done = true
			save(tasks)
			println("Task marked as done:", tasks[idx].Title)
		}
	default:
		println("Unknown command:", command)
	}
}
