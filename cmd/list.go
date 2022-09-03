/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "fmt"
	"fmt"
	"log"
	"os"

	// db "github.com/UtkarshM-hub/todo/DB"
	db "github.com/UtkarshM-hub/todo/DB"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Long: `List all tasks`,
	Run: func(cmd *cobra.Command, args []string) {

		tasks,err:=db.GetTasks()
		if err!=nil{
			log.Fatal("Error fetching data")
		}

		data :=[][]string{}

		for index,item:=range tasks{
			var prio string
			if item.Priority<5{
				prio="low"
			} else if item.Priority>5{
				prio="high"
			} else {
				prio="medium"
			}
			st:="In-Progress"
			if(item.Status){
				st="Completed"
			}
			data = append(data, []string{fmt.Sprint(index+1),item.Task,st,prio})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"#", "Name", "Status","Priority"})

		for _, v := range data {
			table.Append(v)
		}
		table.Render()
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
