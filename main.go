package main

import "fmt"

type Task struct {
	ID int
	Name string
	Done bool
	Priority string
}

func main() {
	fmt.Println("=== Todo List CLI===")

	task := Task{ID: 1, Name: "Belajar GO", Done: false, Priority: "Basic"}
	fmt.Println(task)

	// Printf memakai format verb, jadi \n harus ditulis sendiri.
	// %d -> int, %s -> string, %t -> bool.
	// Jumlah dan urutan argumen harus sesuai dengan jumlah dan urutan verb.
	fmt.Printf("Tugas #%d: %s (selesai: %t, prioritas: %s)\n", task.ID, task.Name, task.Done, task.Priority)
}