package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
	bolt "go.etcd.io/bbolt"
)

const (
	todofile = ".td"
)

// global variables
var (
	WelcomeMsg = `---------Welcome to td App---------
        Here are your todos`
	HomeDir           = GetHomeDir()
	TodosFile         = filepath.Join(HomeDir, todofile)
	defaultTodoBucket = "td"
)

func main() {
	app := &cli.App{
		Name:  "td",
		Usage: "A simple todos management CLI app",
		Action: func(*cli.Context) error {
			fmt.Println(WelcomeMsg)

			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(cCtx *cli.Context) error {
					args := make([]string, cCtx.Args().Len())
					for i := 0; i < cCtx.Args().Len(); i++ {
						args[i] = cCtx.Args().Get(i) // Convert each argument to a string
					}
					newTodo := strings.Join(args, " ")
					return CreateTodo(newTodo)
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("completed task: ", cCtx.Args().First())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
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

// CreateTodo Creates a new todo if that doesn't exist in a new bucket "td" if
// not exists
func CreateTodo(task string) error {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(defaultTodoBucket))
		if err != nil {
			return err
		}

		h := sha256.New()
		h.Write([]byte(task))
		key := h.Sum(nil)
		if err != nil {
			return err
		}
		err = bucket.Put(key, []byte(task))
		if err != nil {
			return err
		}
		// Make progressive count of total todos
		currentCount := bucket.Get([]byte("count"))
		count := 0
		if currentCount != nil {
			count = int(currentCount[0])
		}
		count++
		err = bucket.Put([]byte("count"), []byte{byte(count)})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func ReadTodos(todofile string) {
}

func printTodos() {
}
