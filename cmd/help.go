package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show custom help information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\ntodo help\nSee all available commands")
		fmt.Println("\ntodo add\nAdd a todo")
		fmt.Println("\ntodo add -t \"{task_name}\" -s \"{task_status} (y/n)\"\nAdd a todo")
		fmt.Println("\ntodo list\nView uncompleted todos")
		fmt.Println("\ntodo list -a\nView all todos")
		fmt.Println("\ntodo delete\nDelete a todo")
		fmt.Println("\ntodo complete\nMark a todo as completed")
	},
}
