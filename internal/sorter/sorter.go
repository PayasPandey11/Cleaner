package sorter

import (
	"fmt"
	"log"
	"organiser/internal/util"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
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

	err := util.MoveFileToDirectory(srcPath, destPath)
	if err != nil {
		return false, err
	}
	return true, nil
}

func sortFile(wg *sync.WaitGroup, successCount *uint64, failCount *uint64, basePath string, file os.DirEntry, extensionMap map[string]string) {
	defer wg.Done()
	fileExt := filepath.Ext(file.Name())
	success, err := SortFilesByExtension(basePath, file, fileExt, extensionMap)
	if success {
		atomic.AddUint64(successCount, 1)
		log.Printf("Successfully sorted: %s\n", file.Name())
	} else {
		atomic.AddUint64(failCount, 1)
		log.Printf("Failed to sort file %s: %v", file.Name(), err)
	}
}

func SortFilesConcurrently(basePath string, files []os.DirEntry, extensionMap map[string]string) {
	var wg sync.WaitGroup
	var successCount uint64
	var failCount uint64

	for _, file := range files {
		if !file.IsDir() {
			wg.Add(1)
			go sortFile(&wg, &successCount, &failCount, basePath, file, extensionMap)
		}
	}

	wg.Wait()
	fmt.Printf("Total successful copies: %d\n", successCount)
	fmt.Printf("Total failed operations: %d\n", failCount)
}
