package commands

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"

)

var DeleteCmd = &cobra.Command{
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
        
        for i, task := range Tasks {
            if task.ID == taskID {
                Tasks = append(Tasks[:i], Tasks[i+1:]...)
                fmt.Println("Task deleted successfully.")
                saveTasksToJSON("tasks.json")
                return
            }
        }
        
        fmt.Println("Task not found.")
    },
}

