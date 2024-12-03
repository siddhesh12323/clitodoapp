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

var completeCmd = &cobra.Command{
	Use: "complete",
	Run: func(cmd *cobra.Command, args []string) {

		//----------------------- Check if CSV File exists
		if !fileExists(CSVFILENAME) {
			fmt.Printf("Please run 'todo' command first!")
			return
		}

		ListTasks()

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

		//----------------------- Input to which task to mark complete
		buffreader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter task ID to mark complete:- ")
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

		//----------------------- Update the Task Status
		rows[intid][2] = "Completed" // Assuming column 2 (index 2) stores the status.
		fmt.Printf("Task ID %d marked as complete!\n", intid)

		//----------------------- Write Updated Rows Back to File
		file, writefileerr := os.Create(CSVFILENAME)
		if writefileerr != nil {
			fmt.Println("Error creating file:", writefileerr)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		if writeerr := writer.WriteAll(rows); writeerr != nil {
			fmt.Println("Error writing to file:", writeerr)
			return
		}

		fmt.Println("Task updated successfully!")

	},
}
