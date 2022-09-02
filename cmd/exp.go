/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	db "github.com/UtkarshM-hub/todo/DB"
	"github.com/spf13/cobra"
)

var indexNumber int64

// expCmd represents the exp command
var expCmd = &cobra.Command{
	Use:   "exp",
	Short: "Expand the information of specific task",
	Long: `Expand the information of specific task by providing index`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Expand(indexNumber)
	},
}

func init() {
	rootCmd.AddCommand(expCmd)
	expCmd.Flags().Int64VarP(&indexNumber,"index","i",0,"to specify the index");

	expCmd.MarkFlagRequired("index")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// expCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// expCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
