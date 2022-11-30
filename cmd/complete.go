/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/devaraj-s/terminal-todo/todo"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		item, _ := cmd.Flags().GetInt64("complete")

		if item == -1 {
			fmt.Println("Please enter index to mark complete")
			os.Exit(0)
		}
		todo := &todo.Todos{}
		todo.Load()
		todo.Complete(int(item))
		if err := todo.Store(); err != nil {
			fmt.Println(err)
		}
		todo.List()
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	completeCmd.Flags().Int64P("complete", "c", -1, "Mark a todo complete")
}
