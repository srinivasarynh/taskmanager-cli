package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task id]",
	Short: "delete a task",
	Long:  "delete a task by specifying task id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil || index < 1 {
			fmt.Println("invalid task id")
			return
		}

		tasks := loadTasks()
		if index > len(tasks) {
			fmt.Println("task id out of range")
			return
		}

		tasks = append(tasks[:index-1], tasks[index:]...)
		saveTasks(tasks)
		fmt.Printf("task %d deleted.\n", index)
	},
}
