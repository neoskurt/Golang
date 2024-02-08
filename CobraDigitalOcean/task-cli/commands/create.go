package commands

import (
    "fmt"
    "log"
    "os"
	"encoding/json"

	"github.com/spf13/cobra"
)


var CreateCmd = &cobra.Command{
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
        Tasks = append(Tasks, newTask)
        saveTasksToJSON("tasks.json")
        fmt.Println("New task created successfully.")
    },
	
}

func getLastTaskID() int {
	if len(Tasks) == 0 {
		return 0
	}
	return Tasks[len(Tasks)-1].ID
}

func saveTasksToJSON(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Tasks); err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}
}


