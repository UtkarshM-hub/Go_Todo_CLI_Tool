/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	db "github.com/UtkarshM-hub/todo/DB"
	"github.com/spf13/cobra"
)

var compIndex int64

// compCmd represents the comp command
var compCmd = &cobra.Command{
	Use:   "comp",
	Short: "Set a Task as complete",
	Long: `Set a Task as complete by specifying the index`,
	Run: func(cmd *cobra.Command, args []string) {
		db.MarkAsComplete(compIndex)
	},
}

func init() {
	rootCmd.AddCommand(compCmd)
	compCmd.Flags().Int64VarP(&compIndex,"index","i",0,"specify the index number")
	rootCmd.MarkFlagRequired("index")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
