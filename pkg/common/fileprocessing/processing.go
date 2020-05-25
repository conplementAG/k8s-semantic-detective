package fileprocessing

import (
	"github.com/rs/xid"
	"io/ioutil"
	"os"
	"path/filepath"
)

// WriteStringToTemporaryFile writes the file contents into a file on a temporary disk location
func WriteStringToTemporaryFile(fileContents string, filePath string) string {
	outputFolderPath := createUniqueDirectory()

	generatedFilePath := filepath.Join(outputFolderPath, filePath)
	err := ioutil.WriteFile(generatedFilePath, []byte(fileContents), 0644)
	panicOnError(err)

	return generatedFilePath
}

// DeletePath deletes the file from the disk
func DeletePath(filePath string) {
	err := os.RemoveAll(filePath)
	panicOnError(err)
}

func panicOnError(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func createUniqueDirectory() string {
	folderPath := filepath.Join(".", ".generated", xid.New().String())
	err := os.MkdirAll(folderPath, os.ModePerm)
	panicOnError(err)
	return folderPath
}