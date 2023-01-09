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

func PartitionDataset(dataset map[string][]string, partitionCount int) []map[string][]string {
	// Make a slice of maps with length n
	partitions := make([]map[string][]string, partitionCount)
	for i := range partitions {
		partitions[i] = make(map[string][]string)
	}

	// Iterate over the keys in the data map
	i := 0
	for key, values := range dataset {
		// Iterate over the values in the values slice
		for _, value := range values {
			// Add the key-value pair to the current partition
			partitions[i][key] = append(partitions[i][key], value)
			// Move to the next partition
			i = (i + 1) % partitionCount
		}
	}

	return partitions
}

// TODO: Work on type assertions for Values of Dataset
func TypeAssertDataset(dataset interface{}) map[string][]string {
	asserted, ok := dataset.(map[string]interface{})
	if !ok {
		log.Fatal("FAILURE: Failed to assert dataset to a proper type")
	}

	result := map[string][]string{}
	for key := range asserted {
		assertedVal, _ := (asserted[key]).([]string)
		result[key] = assertedVal
	}
	return result
}
