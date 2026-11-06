package main

import "fmt"

type Task struct {
	ID int
	Name string
	Done bool
	Priority string
}

// listTasks menampilkan semua tugas di slice
func listTasks(tasks []Task){
	if len(tasks) == 0 {
		fmt.Println("Belum ada tugas.")
		return
	}
	
	for _, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d. %s (%s)\n", status, task.ID, task.Name, task.Priority)
	}
}

// addTasks versi pertama: menerima slice biasa
func addTask(tasks *[]Task, name, priority string) []Task {
	newTask := Task{
		ID: len(*tasks) + 1,
		Name: name,
		Done: false,
		Priority: priority,
	}
	return append(*tasks, newTask)
}

func main() {
	fmt.Println("=== Todo List CLI===")

	var tasks []Task

	addTask(&tasks,"Belajar Go", "High")
	addTask(&tasks,"Latihan CLI", "Basic")
	addTask(&tasks,"Latihan Pointer", "High")
	addTask(&tasks,"Latihan Struct", "High")

	listTasks(tasks)
}