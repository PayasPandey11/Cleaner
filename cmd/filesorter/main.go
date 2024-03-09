package main

import (
	"fmt"
	"log"
	"organiser/internal/config"
	"organiser/internal/sorter"
	"organiser/internal/util"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Organising files...")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to get home directory. Error: %v\n", err)
	}
	downloadsPath := filepath.Join(homeDir, "Downloads/testORG")

	fmt.Println("Getting files from: ", downloadsPath)
	contents, err := util.GetAllContentsOfPath(downloadsPath)
	if err != nil {
		log.Fatalf("Unable to get files from %s. Error: %v\n", downloadsPath, err)
	}

	sorter.SortFilesConcurrently(downloadsPath, contents, config.ExtensionMap)
}
