/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/devaraj-s/terminal-todo/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding a todo",
	Long:  `This will add a todo to action`,
	Run: func(cmd *cobra.Command, args []string) {
		item, _ := cmd.Flags().GetString("todo")

		fmt.Println(item)
		todo := &todo.Todos{}
		todo.Load()
		todo.Add(item)
		if err := todo.Store(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringP("todo", "t", "", "Add the todo")
}
