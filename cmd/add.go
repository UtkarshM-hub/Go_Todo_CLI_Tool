/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	database "github.com/UtkarshM-hub/todo/DB"
	"github.com/spf13/cobra"
)

var taskName string
var priority int64
var description string=""

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task to TODO",
	Long: `Add task to TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		database.AddTask(taskName,false,priority,description)
		fmt.Println(" ✅ Task Added Successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&taskName,"name","n","","used to add name")
	addCmd.Flags().Int64VarP(&priority,"priority","p",0,"used to set priority")
	addCmd.Flags().StringVarP(&description,"description","d","","used to set description")

	rootCmd.MarkFlagRequired("name")
	rootCmd.MarkFlagRequired("priority")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
