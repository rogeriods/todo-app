package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/rogeriods/todo-app/todo"
	"github.com/spf13/viper"
)

var priority int

// addCmd respresents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)

		items = append(items, item)
	}

	if err := todo.SaveItems(viper.GetString("datafile"), items); err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}
}

func init() {
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
	// Here you will define your flags and configuration settings

	// Cobra supports Persistent Flags which will work for this command and all subcommands
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle" )
}
