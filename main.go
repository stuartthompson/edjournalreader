package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type JournalEntry struct {
	Timestamp string `json:"timestamp"`
	Event string `json:"event"`
}

func main() {
	file, err := os.Open("test.log")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
	file.Close()
 
	data := JournalEntry{}
	for _, line := range txtlines {
		_ = json.Unmarshal([]byte(line), &data)
		
		fmt.Println("Timestamp: ", data.Timestamp)
		fmt.Println("Event: ", data.Event)
	}
}