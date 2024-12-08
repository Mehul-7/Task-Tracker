/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	// "Task-Tracker-CLI/util"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Gets a list of all",
	Long: `This command will retreive a list of all tasks irrespective of the status`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		// util.Test()
		var gopherName = "dr-who.png"

        if len(args) >= 1 && args[0] != "" {
            gopherName = args[0]
        }

        URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"

        fmt.Println("Try to get '" + gopherName + "' Gopher...")

        // Get the data
        response, err := http.Get(URL)
        if err != nil {
            fmt.Println(err)
        }
        defer response.Body.Close()

        if response.StatusCode == 200 {
            // Create the file
            out, err := os.Create(gopherName + ".png")
            if err != nil {
                fmt.Println(err)
            }
            defer out.Close()

            // Writer the body to file
            _, err = io.Copy(out, response.Body)
            if err != nil {
                fmt.Println(err)
            }

            fmt.Println("Perfect! Just saved in " + out.Name() + "!")
        } else {
            fmt.Println("Error: " + gopherName + " not exists! :-(")
        }
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
