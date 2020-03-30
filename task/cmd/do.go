package cmd

import (
	"errors"
	"fmt"
	todo "github.com/ambye85/gophercises/task/app"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed",
	Long: `Marks a task as completed in the TODO list.

task do 1
You have completed the "clean car" task.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			_, err := strconv.Atoi(args[0])
			if err == nil {
				return nil
			}
		}

		return errors.New("must give a task number")
	},
	Run: doTask,
}

func doTask(cmd *cobra.Command, args []string) {
	task, _ := strconv.Atoi(args[0])
	description, err := todo.Do(task)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("You have completed the \"%s\" task.\n", description)
}

func init() {
	rootCmd.AddCommand(doCmd)
}
