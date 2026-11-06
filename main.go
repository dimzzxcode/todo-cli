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

// listTasks menampilkan semua tugas
func (t *TodoList) List() {
	if len(t.tasks) == 0 {
		fmt.Println("Belum ada tugas.")
		return
	}

	for _, task := range t.tasks {
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

func (t *TodoList) MarkDone(id int) error {
	for i := range t.tasks {
		if t.tasks[i].ID == id {
			t.tasks[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("tugas dengan ID %d tidak ditemukan", id)
}

func (t *TodoList) Delete(id int) error {
	for i, task := range t.tasks {
		if task.ID == id {
			t.tasks = append((t.tasks)[:i], (t.tasks)[i+1:]...)
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
	
	if err := list.MarkDone(1); err != nil {
		fmt.Println("Error:", err)
	}
	
	if err := list.Delete(2); err != nil {
		fmt.Println("Error:", err)
	}

	if err := list.Delete(99); err != nil {
		fmt.Println("Error:", err)
	}
	
	list.Add("Tugas baru", "Basic")

	list.List()
	fmt.Printf("nextID sekarang: %d\n", list.nextID)
}