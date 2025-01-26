package cmd

import (
	"GabrielChaves1/task-tracker/internal/adapters"
	service "GabrielChaves1/task-tracker/internal/services"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateStatus)
}

var updateStatus = &cobra.Command{
	Use:   "update-status",
	Short: "Update the status of a task",
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

		status, err := tasks.IsValidStatus(args[1])

		if err != nil {
			cmd.Println(err)
			return
		}

		result, err := tasks.UpdateStatus(taskID, status)

		if err != nil {
			cmd.Println(err)
			return
		}

		cmd.Println(result)
	},
}
