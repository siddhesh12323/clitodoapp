package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/siddhesh12323/clitodoapp/shared"

	"github.com/spf13/cobra"
)

var (
	t string
	s string
)

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {

		//----------------------- Check if CSV File exists
		if !fileExists(CSVFILENAME) {
			fmt.Printf("Please run 'todo' command first!")
			return
		}

		//----------------------- Check Flag Validity
		s = strings.TrimSpace(s)
		t = strings.TrimSpace(t)
		if s != "n" && s != "N" && s != "Y" && s != "y" {
			fmt.Println("Please enter status as 'y' or 'n' only!")
			return
		}

		//------------- Create and Input a New Task
		var newtask shared.Task
		reader := bufio.NewReader(os.Stdin)
		if t == "" {
			fmt.Print("Enter task name:- ")
			name, _ := reader.ReadString('\n')
			fmt.Println("Enter task status(Enter 'y' for completed and 'n' for not completed):- ")
			status, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			newtask.Name = name
			status = strings.TrimSpace(status)
			if status == "y" || status == "Y" {
				newtask.Status = "Completed"
			} else {
				newtask.Status = "Not Completed"
			}
		} else {
			newtask.Name = t
			if s == "y" || s == "Y" {
				newtask.Status = "Completed"
			} else {
				newtask.Status = "Not Completed"
			}
		}
		newtask.ID = getID()

		//------------------------ Open the CSV File for appending
		file, err := os.OpenFile(CSVFILENAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open file: %s", err)
		}
		defer file.Close()

		//------------------------ Define a CSV Writer
		writer := csv.NewWriter(file)
		defer writer.Flush()

		//------------------------- Append the new task to the CSV File
		row := []string{
			newtask.ID,
			newtask.Name,
			newtask.Status,
			time.Now().Format("2006-01-02 15:04:05"),
		}
		if err := writer.Write(row); err != nil {
			log.Fatalf("Failed to write row: %s", err)
		}

		fmt.Println("Task added!")
	},
}
