package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/icholy/digest"
)

func main() {

	captureCam()

}

func captureCam() {

	cameraName := "cam1"
	url := "http://172.18.3.101/ISAPI/Streaming/Channels/101/picture"
	username := "admin"
	password := "passw0rd"
	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("%s_%s.jpg", cameraName, timestamp)

	client := &http.Client{
		Transport: &digest.Transport{
			Username: username,
			Password: password,
		},
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	res, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Copy the HTTP response body to the file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Println("File downloaded and saved as", fileName)

}
