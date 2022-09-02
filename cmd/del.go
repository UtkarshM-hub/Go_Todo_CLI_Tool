/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	db "github.com/UtkarshM-hub/todo/DB"
	"github.com/spf13/cobra"
)

var index int64

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a task",
	Long: `Delete a task`,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFlags(0)
		if index==0{
			log.Fatal("index not specified")
		}
		db.Delete(index)
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
	delCmd.Flags().Int64VarP(&index,"index","i",0,"to specify and index")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
