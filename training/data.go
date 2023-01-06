package training

import (
	"encoding/base64"
	"log"
	"os"
	"path/filepath"
)

func getImages(imageDirectory string) []string {
	entries, err := os.ReadDir(imageDirectory)
	if err != nil {
		log.Fatal(err)
	}
	// Iterate through all the images and append to images slice
	images := []string{}
	for _, entry := range entries {
		imagePath := filepath.Join(imageDirectory, entry.Name())
		data, err := os.ReadFile(imagePath)
		encoded := base64.StdEncoding.EncodeToString(data)
		if err != nil {
			log.Fatal(err)
		}
		images = append(images, encoded)
	}
	return images
}

func GetDataset(directory string) map[string][]string {
	entries, err := os.ReadDir(directory)

	if err != nil {
		log.Fatal(err)
	}
	result := map[string][]string{}
	for _, entry := range entries {
		projectPath, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fullPath := filepath.Join(projectPath, directory, entry.Name())
		result[entry.Name()] = getImages(fullPath)

	}

	return result
}
