package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple CLI Todo application",
	Long:  "A simple Todo application built in Go using Cobra.",
	Run: func(cmd *cobra.Command, args []string) {

		//------------- Create a csv file if it does not exists
		if !fileExists(CSVFILENAME) {
			file, err := os.Create("tasks.csv")
			if err != nil {
				log.Fatalf("Failed to create file: %s", err)
			}
			defer file.Close()

			// Create a CSV writer
			writer := csv.NewWriter(file)
			defer writer.Flush()

			// Write the header
			header := []string{"ID", "Name", "Status", "Created At"}
			if err := writer.Write(header); err != nil {
				log.Fatalf("Failed to write header: %s", err)
			}
		}

		fmt.Println("Welcome Back to your Todo App! Run 'todo help' to get a list of commands.")
	},
}

func init() {
	addCmd.Flags().StringVarP(&t, "task", "t", "", "Show all tasks, including completed ones")
	addCmd.Flags().StringVarP(&s, "status", "s", "n", "Show all tasks, including completed ones")
	rootCmd.AddCommand(addCmd)
	listCmd.Flags().BoolVarP(&a, "all", "a", false, "Show all tasks, including completed ones")
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.SetHelpFunc(helpCmd.Run)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
