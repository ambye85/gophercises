package cmd

import (
	"fmt"
	todo "github.com/ambye85/gophercises/task/app"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Long: `Add a task to the TODO list:

task add clean car`,
	Args: cobra.MinimumNArgs(1),
	Run:  addTask,
}

func addTask(cmd *cobra.Command, args []string) {
	task := strings.Join(args, " ")
	err := todo.Add(task)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Added \"%s\" to your task list.\n", task)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
