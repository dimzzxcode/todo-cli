package main

import "fmt"

type Task struct {
	ID       int
	Name     string
	Done     bool
	Priority string
}

// listTasks menampilkan semua tugas di slice
func listTasks(tasks []Task) {
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
func addTask(tasks *[]Task, name, priority string) {
	newTask := Task{
		ID:       len(*tasks) + 1,
		Name:     name,
		Done:     false,
		Priority: priority,
	}
	*tasks = append(*tasks, newTask)
}

func markDone(tasks []Task, id int) error {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("tugas dengan ID %d tidak ditemukan", id)
}

func deleteTask(tasks *[]Task, id int) error {
	for i, task := range *tasks {
		if task.ID == id {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("tugas dengan ID %d tidak ditemukan", id)
}

func main() {
	fmt.Println("=== Todo List CLI===")

	var tasks []Task

	addTask(&tasks, "Belajar Go", "High")
	addTask(&tasks, "Latihan CLI", "Basic")
	addTask(&tasks, "Latihan Pointer", "High")
	addTask(&tasks, "Latihan Struct", "High")
	addTask(&tasks, "Tugas baru", "Basic")

	if err := markDone(tasks, 1); err != nil {
		fmt.Println("Error:", err)
	}

	if err := deleteTask(&tasks, 2); err != nil {
		fmt.Print("Error:", err)
	}

	if err := markDone(tasks, 99); err != nil {
		fmt.Println("Error:", err)
	}

	listTasks(tasks)
}