package cmd

import (
	"Task-Tracker-CLI/util"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <Task Description>",
	Short: "Add a new task",
	Long: `This command will be used to create a new task and add it to the list of exisitng tasks`,
	Run: func(cmd *cobra.Command, args[] string){
		taskObj := util.New(args[0])
		util.AddTask(taskObj)
	},
}

func init(){
	rootCmd.AddCommand(addCmd)
}