package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var (
	a bool
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		ListTasks()
	},
}

func ListTasks() {
	//----------------------- Check if CSV File exists
	if !fileExists(CSVFILENAME) {
		fmt.Printf("Please run 'todo' command for first time!")
		return
	}

	//----------------------- Open the file
	file, err := os.Open(CSVFILENAME)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	//----------------------- Read all the rows
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	//----------------------- Print out the list
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 6, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tStatus\tCreated")
	for i, row := range rows {
		if i == 0 {
			continue
		}
		var status bool
		if row[2] == "Completed" {
			status = true
		} else {
			status = false
		}
		if a || !status {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", i, row[1], row[2], getCreatedAgo(row[3]))
		}
	}

	w.Flush()
}

func ListAllTasks() {
	//----------------------- Check if CSV File exists
	if !fileExists(CSVFILENAME) {
		fmt.Printf("Please run 'todo' command for first time!")
		return
	}

	//----------------------- Open the file
	file, err := os.Open(CSVFILENAME)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	//----------------------- Read all the rows
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	//----------------------- Print out the list
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 6, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tStatus\tCreated")
	for i, row := range rows {
		if i == 0 {
			continue
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", i, row[1], row[2], getCreatedAgo(row[3]))
	}

	w.Flush()
}
