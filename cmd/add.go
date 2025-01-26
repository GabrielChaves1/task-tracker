package cmd

import (
	"GabrielChaves1/task-tracker/internal/adapters"
	service "GabrielChaves1/task-tracker/internal/services"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(add)
}

var add = &cobra.Command{
	Use:   "add",
	Short: "Create a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		description := strings.Join(args, " ")

		tasks := service.NewTaskService(
			adapters.NewJSONStorage(),
		)

		result, err := tasks.Add(description)

		if err != nil {
			cmd.Println(err)
			return
		}

		cmd.Println(result)
	},
}
