/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Task-Tracker-CLI/util"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list \n  todo\n  in-progress\n  done",
	Short: "Gets a list of all",
	Long: `This command will retreive a list of all tasks irrespective of the status`,
	Run: func(cmd *cobra.Command, args []string) {
		var status string

        if len(args) > 0 {
            status = args[0]
        }

        util.ListTasks(status)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
