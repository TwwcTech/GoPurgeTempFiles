package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func purgeDownloadsFolder() ([]string, error) {
	// get user's Downloads folder
	usrHomePath, err := os.UserHomeDir()
	if err != nil {
		_ = fmt.Errorf("error getting user home directory: %v", err)
	}
	downloadsPath := filepath.Join(usrHomePath, "Downloads")

	// initialize empty slice
	var downloadedFiles []string
	// walk the Downloads folder
	fileErr := filepath.Walk(downloadsPath, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			_ = fmt.Errorf("error walking the path %q: %v", path, err)
		}
		if !fileInfo.IsDir() {
			// if the items is not a directory append the file to the slice
			downloadedFiles = append(downloadedFiles, path)
		}
		return nil

	})
	if fileErr != nil {
		_ = fmt.Errorf("error walking the path %q: %v", downloadsPath, fileErr)
	}
	return downloadedFiles, nil
}

func purgeTempFiles() (int, error) {
	tmpPath := os.TempDir()

	var tempFiles []string
	tmpFileErr := filepath.Walk(tmpPath, func(path string, d os.FileInfo, err error) error {
		if err != nil {
			_ = fmt.Errorf("error walking the path %q: %v", tmpPath, err)
		}
		if !d.IsDir() {
			tempFiles = append(tempFiles, path)
		}
		return nil
	})

	var numberOfFileRemoved []int
	for i, file := range tempFiles {
		err := os.Remove(file)
		if err != nil {
			_ = fmt.Errorf("%d: file in use: %q", i+1, file)
			continue
		} else {
			numberOfFileRemoved = append(numberOfFileRemoved, i+1)
		}
	}

	return len(numberOfFileRemoved), tmpFileErr
}

func main() {

}
