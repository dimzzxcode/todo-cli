package main

import "fmt"

type Task struct {
	ID       int
	Name     string
	Done     bool
	Priority string
}

type TodoList struct {
	tasks []Task
	nextID int // id naik tidak turun
}

func NewTodoList() *TodoList {
	return &TodoList{nextID: 1}
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

func (t *TodoList) Add(name, priority string) {
	task := Task{
		ID:       t.nextID,
		Name:     name,
		Done:     false,
		Priority: priority,
	}
	t.tasks = append(t.tasks, task)
	t.nextID++
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

	list := NewTodoList()

	list.Add("Belajar Go", "High")
	list.Add("Latihan CLI", "Basic")
	list.Add("Latihan Pointer", "High")
	list.Add("Latihan Struct", "High")
	list.Add("Tugas baru", "Basic")

	fmt.Printf("%+v\n", list.tasks)
	fmt.Printf("nextID sekarang: %d\n", list.nextID)
}