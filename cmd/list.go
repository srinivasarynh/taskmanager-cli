package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Long:  "display all tasks with their status",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		if len(tasks) == 0 {
			fmt.Println("no tasks found")
			return
		}

		for i, task := range tasks {
			status := "Pending"
			if task.Done {
				status = "Done"
			}
			fmt.Printf("%d. %s [%s]\n", i+1, task.Description, status)
		}
	},
}
