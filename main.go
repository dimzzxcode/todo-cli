package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

// List menampilkan semua tugas.
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

func readInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func showMenu(){
	fmt.Println()
	fmt.Println("=== Todo List CLI===")
	fmt.Println("1. Lihat Tugas")
	fmt.Println("2. Tambah Tugas")
	fmt.Println("3. Tandai Tugas")
	fmt.Println("4. Hapus Tugas")
	fmt.Println("0. Keluar")
}

func main() {
	list := NewTodoList()
	reader := bufio.NewReader(os.Stdin)

	for {
		showMenu()
		choice := readInput(reader,"Pilih menu: ")

		switch choice {
		case "1":
			list.List()
		case "2":
			name := readInput(reader, "Nama tugas: ")
			if name == "" {
				fmt.Println("Nama tugas tidak boleh kosong.")
				continue
			}
			priority := readInput(reader, "Prioritas (High/Basic): ")
			list.Add(name, priority)
			fmt.Println("Tugas ditambahkan.")
		case "3":
			id, err := strconv.Atoi(readInput(reader, "ID tugas: "))
			if err != nil {
				fmt.Println("ID harus berupa angka.")
				continue
			}
			if err := list.MarkDone(id); err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("Tugas ditandai selesai.")
		case "4":
			id, err := strconv.Atoi(readInput(reader, "ID tugas: "))
			if err != nil {
				fmt.Println("ID harus berupa angka.")
				continue
			}
			if err := list.Delete(id); err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("Tugas dihapus.")
		case "0":
			fmt.Println("Sampai jumpa!")
			return

		default:
			fmt.Println("Menu tidak dikenal.")
		}
	}
}