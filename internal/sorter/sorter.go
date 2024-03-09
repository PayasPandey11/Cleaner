package sorter

import (
	"fmt"
	"organiser/internal/util"
	"os"
	"path/filepath"
	"strings"
)

func SortFilesByExtension(basePath string, file os.DirEntry, fileExt string, extensionMap map[string]string) (bool, error) {
	// Convert the file extension to lowercase before looking it up in the map
	lowercaseExt := strings.ToLower(fileExt)
	targetDir, exists := extensionMap[lowercaseExt]
	if !exists {
		return false, fmt.Errorf("no target directory for file extension: %s", fileExt)
	}

	// Construct the source and destination paths
	srcPath := filepath.Join(basePath, file.Name())
	destPath := filepath.Join(basePath, targetDir, file.Name())

	// Ensure the target directory exists
	os.MkdirAll(filepath.Join(basePath, targetDir), os.ModePerm)

	// Copy the file
	err := util.MoveFileToDirectory(srcPath, destPath)
	if err != nil {
		return false, err
	}
	return true, nil
}
