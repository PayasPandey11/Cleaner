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
		log.Printf("Unable to get home directory. Error: %v\n", err)
		return
	}
	downloadsPath := filepath.Join(homeDir, "Downloads/testORG")

	fmt.Println("getting files from: ", downloadsPath)
	contents, err := util.GetAllContentsOfPath(downloadsPath)
	if err != nil {
		log.Printf("Unable to get files from %s. Error: %v\n", downloadsPath, err)
		return
	}

	successCount := 0
	failCount := 0

	for _, content := range contents {
		if !content.IsDir() {
			fileExt := filepath.Ext(content.Name())
			success, err := sorter.SortFilesByExtension(downloadsPath, content, fileExt, config.ExtensionMap)
			if success {
				successCount++
			} else {
				log.Printf("Failed to sort file %s: %v", content.Name(), err)
				failCount++
			}
		}
	}

	fmt.Printf("Total successful copies: %d\n", successCount)
	fmt.Printf("Total failed operations: %d\n", failCount)
}
