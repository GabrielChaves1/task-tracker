package cmd

import (
	"GabrielChaves1/task-tracker/internal/adapters"
	service "GabrielChaves1/task-tracker/internal/services"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(remove)
}

var remove = &cobra.Command{
	Use:   "remove",
	Short: "Remove a task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		taskID, err := strconv.Atoi(args[0])

		if err != nil {
			cmd.Println(err)
			return
		}

		tasks := service.NewTaskService(
			adapters.NewJSONStorage(),
		)

		result, err := tasks.Remove(taskID)

		if err != nil {
			cmd.Println(err)
			return
		}

		cmd.Println(result)
	},
}
