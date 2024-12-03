package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/mergestat/timediff"
)

var CSVFILENAME string = "tasks.csv"

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	// Check if the error is "file does not exist"
	return !os.IsNotExist(err)
}

func getID() string {
	taskID := uuid.New().String()
	return taskID
}

func getCreatedAgo(createdAt string) string {
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", createdAt, time.Local)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return "error!"
	}
	str := timediff.TimeDiff(parsedTime)
	return str
}
