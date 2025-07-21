package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [task number]",
	Short: "mark a task as done",
	Long:  "mark a task as done by specifying task id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil || index < 1 {
			fmt.Println("invalid task id")
			return
		}

		tasks := loadTasks()
		if index > len(tasks) {
			fmt.Println("task id out of range.")
			return
		}

		tasks[index-1].Done = true
		saveTasks(tasks)
		fmt.Printf("task %d marked as done.\n", index)
	},
}
