package commands

import (
    "fmt"
    "log"
    "strconv"

	"github.com/spf13/cobra"

)




var PrintCmd = &cobra.Command{
    Use:   "print",
    Short: "Print tasks",
    Long:  "Print all tasks or print a specific task by ID.",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 0 {
            fmt.Println("Afficher tout les taches")
            PrintAllTasks()
        } else {
            // Print specific task by ID
            taskID, err := strconv.Atoi(args[0])
            if err != nil {
                log.Fatalf("ID inconnu: %v", err)
            }
            fmt.Printf("Afficher les taches par ID %d...\n", taskID)
            PrintTaskByID(taskID)
        }
    },
}

func PrintAllTasks() {
    for _, task := range Tasks {
        fmt.Printf("ID: %d, Name: %s, Deadline: %s\n", task.ID, task.Name, task.Deadline)
    }
}

func PrintTaskByID(id int) {
    for _, task := range Tasks {
        if task.ID == id {
            fmt.Printf("ID: %d, Name: %s, Deadline: %s\n", task.ID, task.Name, task.Deadline)
            return
        }
    }
    fmt.Println("La tache n'existe pas")
}
