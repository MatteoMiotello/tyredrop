package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"pillowww/titw/internal/bootstrap"
	"pillowww/titw/pkg/utils"
	"strings"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitDb()
}

func main() {
	records, err := utils.CsvReadFile("./storage/tmp/pnd_tires_items.csv")
	if err != nil {
		panic(err)
	}

	for index, record := range records {
		if index < 1 {
			continue
		}

		slices := strings.Split(record[0], ";")

		if len(slices) < 25 {
			continue
		}

		imageUrl := strings.Trim(slices[25], "\"")
		fileName := strings.Trim(slices[3], "\"")

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{Transport: tr}
		response, err := client.Get(imageUrl)
		if err != nil {
			fmt.Println(err)
			continue
		}

		contentType := response.Header.Get("Content-Type")
		extensionsByType, err := mime.ExtensionsByType(contentType)

		if err != nil {
			fmt.Println(err)
			continue
		}

		extension := new(string)

		if extensionsByType == nil {
			extension = nil
		}

		if len(extensionsByType) > 0 {
			extension = &extensionsByType[len(extensionsByType)-1]
		}

		if extension == nil || (*extension != ".png" && *extension != ".jpg" && *extension != ".jpeg") {
			fmt.Println("incompatible_image")
			continue
		}

		filename := fileName + *extension
		filePath := "./storage/images/" + filename

		if err != nil {
			fmt.Println(err)
			continue
		}

		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)

		_, err = io.Copy(file, response.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = file.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}

	}
}
