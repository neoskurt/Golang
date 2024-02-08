package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Deadline string `json:"deadline"`
}

var tasks []Task

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "A CLI tool to manage tasks",
	Long:  "A CLI tool to manage tasks. Use it to print all tasks or print a specific task by ID.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Task CLI. Use --help for usage instructions.")
	},
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print tasks",
	Long:  "Print all tasks or print a specific task by ID.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Afficher tout les taches")
			printAllTasks()
		} else {
			// Print specific task by ID
			taskID, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("ID inconnu: %v", err)
			}
			fmt.Printf("Afficher les taches par ID %d...\n", taskID)
			printTaskByID(taskID)
		}
	},
}

func printAllTasks() {
	for _, task := range tasks {
		fmt.Printf("ID: %d, Name: %s, Deadline: %s\n", task.ID, task.Name, task.Deadline)
	}
}

func printTaskByID(id int) {
	for _, task := range tasks {
		if task.ID == id {
			fmt.Printf("ID: %d, Name: %s, Deadline: %s\n", task.ID, task.Name, task.Deadline)
			return
		}
	}
	fmt.Println("La tache n'existe pas")
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Long:  "Create a new task and save it to the task list.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal("Usage: task-cli create <Name> <Deadline>")
		}

		lastID := getLastTaskID()
		newID := lastID + 1

		newTask := Task{
			ID:       newID,
			Name:     args[0],
			Deadline: args[1],
		}
		tasks = append(tasks, newTask)
		saveTasksToJSON("tasks.json")
		fmt.Println("New task created successfully.")
	},
}

func getLastTaskID() int {
	if len(tasks) == 0 {
		return 0
	}
	return tasks[len(tasks)-1].ID
}


func saveTasksToJSON(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tasks); err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}
}

var updateCmd = &cobra.Command{
    Use:   "update",
    Short: "Update a task by ID",
    Long:  "Update a task by its ID with new name and deadline.",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) != 3 {
            log.Fatal("Usage: task-cli update <ID> <Name> <Deadline>")
        }
        
        taskID, err := strconv.Atoi(args[0])
        if err != nil {
            log.Fatalf("Invalid ID: %v", err)
        }
        
        updatedName := args[1]
        updatedDeadline := args[2]
        
        for i, task := range tasks {
            if task.ID == taskID {
                tasks[i].Name = updatedName
                tasks[i].Deadline = updatedDeadline
                fmt.Println("Task updated successfully.")
                saveTasksToJSON("tasks.json")
                return
            }
        }
        
        fmt.Println("Task not found.")
    },
}

var deleteCmd = &cobra.Command{
    Use:   "delete",
    Short: "Delete a task by ID",
    Long:  "Delete a task by its ID.",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) != 1 {
            log.Fatal("Usage: task-cli delete <ID>")
        }
        
        taskID, err := strconv.Atoi(args[0])
        if err != nil {
            log.Fatalf("Invalid ID: %v", err)
        }
        
        for i, task := range tasks {
            if task.ID == taskID {
                tasks = append(tasks[:i], tasks[i+1:]...)
                fmt.Println("Task deleted successfully.")
                saveTasksToJSON("tasks.json")
                return
            }
        }
        
        fmt.Println("Task not found.")
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
    rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(printCmd)
	rootCmd.AddCommand(createCmd)
}

func main() {
	loadTasksFromJSON("tasks.json") 
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadTasksFromJSON(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
}