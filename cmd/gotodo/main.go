package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	todofile = ".gotodo"
)

// todo definition
type Todo struct {
	Task   string
	Status bool
	Id     int
}

// global variables
var (
	WelcomeMsg = `
---------Welcome to GoTodo App---------
          Here are your todos
`
	HomeDir   = GetHomeDir()
	TodosFile = filepath.Join(HomeDir, todofile)
	Todos     = []Todo{}
)

func main() {
	fmt.Println(WelcomeMsg)
	_, err := os.Stat(TodosFile)
	if os.IsNotExist(err) {
		fmt.Println(TodosFile, "File does not exist")
		fmt.Println("Creating .gotodo file in", HomeDir)
		CreateTodoFile()
	} else if err != nil {
		fmt.Println("Error:", err)
	}
}

// GetHomeDir returns a home directory
func GetHomeDir() string {
	HomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error:", err)
	}
	return HomeDir
}

// CreateTodoFile Creates a todofile if that doesn't exist
func CreateTodoFile() {
	file, err := os.Create(TodosFile)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println("Successfully created file", TodosFile)
	defer file.Close()
}

func ReadTodos(todofile string) {
}

func printTodos() {
}
