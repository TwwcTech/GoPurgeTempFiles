package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func purgeDownloadsFolder() error {
	usrHomePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error getting user home directory: %v", err)
	}
	downloadsPath := filepath.Join(usrHomePath, "Downloads")

	var downloadedFiles []string
	fileErr := filepath.Walk(downloadsPath, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error walking the path %q: %v", path, err)
		}
		if !fileInfo.IsDir() {
			downloadedFiles = append(downloadedFiles, path)
		}

		return nil
	})

	var numberOfDownloadFilesRemoved []int
	for i, files := range downloadedFiles {
		err := os.Remove(files)
		if err != nil {
			fmt.Printf("%d: %v", i+1, err.Error())
			continue
		}
		numberOfDownloadFilesRemoved = append(numberOfDownloadFilesRemoved, i+1)
	}

	if fileErr != nil {
		fmt.Printf("error walking the path %q: %v", downloadsPath, fileErr)
	}

	fmt.Printf("number of download files removed: %d", len(numberOfDownloadFilesRemoved))
	return nil
}

func purgeTempFiles() error {
	tmpPath := os.TempDir()

	var tempFiles []string
	tmpFileErr := filepath.Walk(tmpPath, func(path string, d os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error walking the path %q: %v", tmpPath, err)
		}
		if !d.IsDir() {
			tempFiles = append(tempFiles, path)
		}

		return nil
	})

	var numberOfFilesRemoved []int
	for i, file := range tempFiles {
		err := os.Remove(file)
		if err != nil {
			fmt.Printf("%d: %v\n", i+1, err.Error())
			continue
		}
		numberOfFilesRemoved = append(numberOfFilesRemoved, i+1)
	}

	if tmpFileErr != nil {
		fmt.Printf("error walking the path %q: %v", tmpPath, tmpFileErr)
	}

	fmt.Printf("\nnumber of temp files removed: %d", len(numberOfFilesRemoved))
	return nil
}

func main() {
	fmt.Println("Go Purge Temp Files v.1.0")
	fmt.Println("-----------------------------")

	fmt.Println("\nRemoving files in the downloads folder...")
	rmvDownloadsErr := purgeDownloadsFolder()
	if rmvDownloadsErr != nil {
		_ = fmt.Errorf("unable to remove files: %w", rmvDownloadsErr)
	}

	fmt.Println("\n\nRemoving temporary system files...")
	rmvTempErr := purgeTempFiles()
	if rmvTempErr != nil {
		_ = fmt.Errorf("unable to remove temp files: %w", rmvTempErr)
	}
	fmt.Println("\n--------------------------------")

	fmt.Println("\n\nPress the 'Enter' key to exit")
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		return
	}
}
