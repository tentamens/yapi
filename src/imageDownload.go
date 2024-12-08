package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

var downloadedImages = map[string]string{}

func downloadAllImages(data ServerResponse) {
	return
	for _, element := range data.Hits {

		downloadImage(element.Icon, "./tmpImage")
		downloadedImages[element.Name] = translateImage()
	}
	fmt.Print(downloadedImages["Discord"])
}

func downloadImage(url string, filePath string) error {
	// Send a GET request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download image: status code %d", resp.StatusCode)
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the image data to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func translateImage() string {
	cmd := exec.Command("./pixelizer", "./tmpImage", "3", "3", "3", "print")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("hi")
		panic(err)
	}

	return string(output)
}
