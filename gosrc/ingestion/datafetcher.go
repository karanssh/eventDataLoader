package ingestion

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

// fetchFilesFromFolder fetches list of files from a folder
func fetchFilesFromFolder(folderName string) []fs.DirEntry {
	files, err := os.ReadDir(folderName)
	if err != nil {
		log.Fatalf("fetchFilesFromFolder() err: %v", err)
	}
	return files
}

// readFilesFetchedFromFolder reads files
// unused
func readFilesFetchedFromFolder() {
	filesList := fetchFilesFromFolder("")
	for _, fs := range filesList {
		if !fs.IsDir() { //ignore folders
			info, err := fs.Info()
			if err != nil {
				log.Fatalf("fs.Info() err: %v", err)
			}
			fileInfo := fmt.Sprintf("fileName %s size: %s", fs.Name(), ByteCountSI(info.Size()))
			log.Print(fileInfo)
		}
	}
}

// ByteCountSI converts bytes to human readable format
// thanks https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
