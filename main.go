package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var client = http.Client{}

func downloadFile(url string) error {
	sp := strings.Split(url, "/")
	var newg string
	if strings.HasPrefix(url, "github.com") || strings.HasPrefix(url, "https://github.com") {
		newg = url + "/archive/refs/heads/master.zip"
	} else {
		log.Fatal("Unsupported domain given.")
	}
	err := os.Mkdir(sp[len(sp)-1], 1)
	if err != nil {
		return err
	}
	file, err := os.Create(filepath.Join(sp[len(sp)-1], "main.zip"))
	if err != nil {
		return err
	}
	g, err := client.Get(newg)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, g.Body)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	fmt.Println("Attempting to download file...")
	err := downloadFile(os.Args[1])
	if err != nil {
		log.Fatal("Failed to download file/folder: ", err)
	}
	fmt.Println("Successfully downloaded repository zip file!")
	link := strings.Split(os.Args[1], "/")
	zip, err := zip.OpenReader(filepath.Join(link[len(link)-1], "/main.zip"))
	if err != nil {
		log.Fatal("Failed to unpack file/folder: ", err)
	}
	defer zip.Close()
	for _, fi := range zip.File {
		fmt.Printf("Unpacking %s..", fi.Name)
		f, err := os.Create(fi.Name)
		if err != nil {
			log.Fatal("Failed to unpack file/folder: ", err)
		}
		readerr, err := fi.Open()
		if err != nil {
			log.Fatal("Failed to unpack file/folder: ", err)
		}
		io.Copy(f, readerr)
		fmt.Print(" [Done]\n")
	}
	fmt.Println("Successfully unzipped files, Exit Code 0")
	os.Exit(0)
}
