package util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copyFileToDirectory(srcPath string, destPath string) error {
	sourceFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}
	return destinationFile.Sync()
}

func MoveFileToDirectory(srcPath string, destPath string) error {
	if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
		return err
	}

	if err := os.Rename(srcPath, destPath); err != nil {
		return err
	}
	return nil
}

func GetAllContentsOfPath(path string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	fmt.Println("files: ", len(files))
	return files, nil
}
