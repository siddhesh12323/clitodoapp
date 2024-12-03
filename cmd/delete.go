package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete",
	Run: func(cmd *cobra.Command, args []string) {
		//----------------------- Check if CSV File exists
		if !fileExists(CSVFILENAME) {
			fmt.Printf("Please run 'todo' command first!")
			return
		}

		ListAllTasks()

		//----------------------- Open the file
		file, openfileerr := os.Open(CSVFILENAME)
		if openfileerr != nil {
			fmt.Println("Error opening file:", openfileerr)
			return
		}
		defer file.Close()

		//----------------------- Read all Rows
		reader := csv.NewReader(file)
		rows, filereaderr := reader.ReadAll()
		if filereaderr != nil {
			fmt.Println("Error reading file:", filereaderr)
			return
		}

		//----------------------- Input to which task to delete
		buffreader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter task ID to delete:- ")
		id, _ := buffreader.ReadString('\n')
		trimmedid := strings.TrimSpace(id)
		intid, intiderr := strconv.Atoi(trimmedid)
		if intiderr != nil {
			fmt.Printf("Invalid ID! Please enter a number!\nError:-  %s\n", intiderr)
			return
		}
		if intid > len(rows)-1 || intid < 0 {
			fmt.Println("ID doesn't exist!")
			return
		}

		// Filter out the row to remove
		var updatedRows [][]string
		for i, row := range rows {
			if i == 0 || i != intid { // Keep the header and non-matching rows
				updatedRows = append(updatedRows, row)
			}
		}

		// Write the updated rows back to the file
		file.Close() // Close the file before rewriting
		file, createfileerr := os.Create(CSVFILENAME)
		if createfileerr != nil {
			fmt.Println("Error creating file:", createfileerr)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		err := writer.WriteAll(updatedRows)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("Task deleted! IDs are now reassigned, Run 'todo list' to check!")

	},
}
