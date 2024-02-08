package main


import (
    "log"
    "os"
	"encoding/json"

	"github.com/spf13/cobra"
	"CobraDigitalOcean/task-cli/commands"

)

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "A CLI tool to manage tasks",
	Long:  "A CLI tool to manage tasks. Use it to print all tasks or print a specific task by ID.",
	
}

func init() {
    rootCmd.AddCommand(commands.DeleteCmd)
    rootCmd.AddCommand(commands.UpdateCmd)
	rootCmd.AddCommand(commands.PrintCmd)
	rootCmd.AddCommand(commands.CreateCmd)
}


func main() {
    loadTasksFromJSON("tasks.json")

    if err := rootCmd.Execute(); err != nil {
        log.Println(err)
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
	if err := decoder.Decode(&commands.Tasks); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
}




