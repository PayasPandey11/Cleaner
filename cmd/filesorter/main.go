package main

import (
	"fmt"
	"log"
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

	extensionMap := map[string]string{
		".pdf":  "Documents",
		".docx": "Documents",
		".doc":  "Documents",
		".txt":  "Documents",
		".webp": "Images",
		".heic": "Images",
		".gif":  "Images",
		".svg":  "Images",
		".png":  "Images",
		".jpg":  "Images",
		".jpeg": "Images",
		".gz":   "Archives",
		".zip":  "Archives",
		".rar":  "Archives",
		".jar":  "Archives",
		".csv":  "Spreadsheets",
		".xlsx": "Spreadsheets",
		".xls":  "Spreadsheets",
		".mp4":  "Media",
		".mp3":  "Media",
		".mov":  "Media",
		".webm": "Media",
		".exe":  "Installers",
		".pkg":  "Installers",
		".dmg":  "Installers",
		".html": "Code",
		".css":  "Code",
		".xml":  "Code",
		".js":   "Code",
		".json": "Code",
		".py":   "Code",
	}

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
			success, err := sorter.SortFilesByExtension(downloadsPath, content, fileExt, extensionMap)
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
