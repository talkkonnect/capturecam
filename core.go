package capturecam

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/icholy/digest"
)

func Init(configFile string) {

	log.Printf("info: %s\n", configFile)
	cameraName := "cam1"
	urlBase := "http://172.18.3.101"
	urlSuffix := "/ISAPI/Streaming/Channels/101/picture"
	url := urlBase + urlSuffix
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
