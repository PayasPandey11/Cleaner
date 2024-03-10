package main

import (
	"fmt"
	"log"
	"organiser/internal/config"
	"organiser/internal/sorter"
	"organiser/internal/util"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Please select a directory.")

	// AppleScript command to open Finder and ask the user to choose a folder
	appleScript := `osascript -e 'tell application "Finder" to return POSIX path of (choose folder with prompt "Select the folder:")'`

	// Execute the AppleScript command
	cmd := exec.Command("bash", "-c", appleScript)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	selectedPath := strings.TrimSpace(string(output))

	fmt.Println("You selected:", selectedPath)

	contents, err := util.GetAllContentsOfPath(selectedPath)

	if err != nil {
		log.Fatalf("Unable to get files from %s. Error: %v\n", selectedPath, err)
	}

	sorter.SortFilesConcurrently(selectedPath, contents, config.ExtensionMap)
}
