package cmd

import (
	"fmt"
	todo "github.com/ambye85/gophercises/task/app"
	"github.com/spf13/cobra"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long: `List the tasks in the TODO list.

task list
You have the following tasks:
1. clean car
2. mow lawn`,
	Args: cobra.NoArgs,
	Run:  listTasks,
}

func listTasks(cmd *cobra.Command, args []string) {
	fmt.Println("You have the following tasks:")
	list, err := todo.List()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i, task := range list {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
