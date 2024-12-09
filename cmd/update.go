package cmd

import (
	"Task-Tracker-CLI/util"
	"log"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing task",
	Long:  `This command will be used to update an existing task	through its UID`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			log.Fatal("missing arguments!")
		}

		util.UpdateTask(args[0], args[2], args[1])
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}