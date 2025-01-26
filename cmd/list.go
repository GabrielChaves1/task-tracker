package cmd

import (
	"GabrielChaves1/task-tracker/internal/adapters"
	service "GabrielChaves1/task-tracker/internal/services"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(list)
}

var list = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := service.NewTaskService(
			adapters.NewJSONStorage(),
		)

		var status string
		if len(args) > 0 {
			status = args[0]
		} else {
			status = ""
		}

		status, err := tasks.IsValidStatus(status)

		if err != nil {
			cmd.Println(err)
			return
		}

		result, err := tasks.List(status)

		if err != nil {
			cmd.Println(err)
			return
		}

		cmd.Println(result)
	},
}
