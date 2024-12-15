package cmd

import (
	"Task-Tracker-CLI/util"
	"log"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <Task UID>",
	Short: "Deletes the task with the given id.",
	Long:  `This command will delete the task with the given id.`,
	Run: func(cmd *cobra.Command, args []string) {
		var status string

		if len(args) == 0 {
			log.Fatal("Missing arguments!")
		}
		if len(args) > 1 {
			log.Fatal("Too many arguments!")
		}

		status = args[0]

		util.DeleteTask(status)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
