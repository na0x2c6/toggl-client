package main

import (
	"fmt"
	"log"
	"os"

	toggl "github.com/na0x2c6/toggl-client"
)

func main() {
	client := toggl.New(os.Getenv("TOGGL_TOKEN"), "api_token")

	entries, err := client.GetTimeEntries(false, false, nil, nil, nil, nil)
	if err != nil {
		log.Fatalf("Error getting time entries: %v", err)
	}

	fmt.Printf("Found %d time entries\n", len(entries))
	for _, entry := range entries {
		fmt.Printf("- %s (%s)\n", entry.Description, entry.Start)
	}
}
