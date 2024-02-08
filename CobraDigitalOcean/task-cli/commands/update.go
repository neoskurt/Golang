package commands

import (
    "fmt"
    "log"
    "strconv"

	"github.com/spf13/cobra"

)

var UpdateCmd = &cobra.Command{
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
        
        for i, Task := range Tasks {
            if Task.ID == taskID {
                Tasks[i].Name = updatedName
                Tasks[i].Deadline = updatedDeadline
                fmt.Println("Task updated successfully.")
                saveTasksToJSON("tasks.json")
                return
            }
        }
        
        fmt.Println("Task not found.")
    },
}
