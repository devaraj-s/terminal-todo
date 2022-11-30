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

// delegateCmd represents the delegate command
var delegateCmd = &cobra.Command{
	Use:   "delegate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		delegateTo, _ := cmd.Flags().GetString("delegate")
		index, _ := cmd.Flags().GetInt("task")

		if index == -1 {
			fmt.Println("Please enter the index of the task")
			os.Exit(0)
		}
		if delegateTo == "" {
			fmt.Println("Please enter the person to delegate to")
			os.Exit(0)
		}
		todo := &todo.Todos{}
		todo.Load()
		todo.Delegate(index, delegateTo)
		if err := todo.Store(); err != nil {
			fmt.Println(err)
		}
		todo.List()
	},
}

func init() {
	rootCmd.AddCommand(delegateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delegateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	delegateCmd.Flags().StringP("delegate", "d", "", "Delegate task to")
	delegateCmd.Flags().IntP("task", "t", -1, "Task id")
}
