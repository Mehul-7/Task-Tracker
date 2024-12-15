/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Task-Tracker-CLI/util"
	"log"

	"github.com/spf13/cobra"
)

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress <Task UID>",
	Short: "Change status of task to in-progress",
	Long: `A longer description that spans multiple lines and likely contains examples and usage of using your command. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Invalid no. of arguments")
		}

		util.UpdateStatus(args[0], "in-progress")
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markInProgressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markInProgressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
