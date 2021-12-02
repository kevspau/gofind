package src

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var client = http.Client{}

func downloadFile_github(url string) (string, error) {
	sp := strings.Split(url, "/")
	primeBranch := false
	for _, v := range sp {
		if v == "tree" {
			primeBranch = true
		}
	}
	if primeBranch {
		req, e := http.Get(url + "/archive/refs/heads/master.zip")
		if e != nil {
			fmt.Println("Error retrieving content")
			return "", e
		}
		file, e := os.Create(sp[len(sp)-1])
		if e != nil {
			fmt.Println("Error creating file")
			return "", e
		}
		io.Copy(file, req.Body)
		return sp[len(sp)-1], nil
	} else if !primeBranch {
		x := strings.Replace(url, "/tree", "", -1)
		x = strings.Replace(x, "/"+sp[len(sp)-1], "", -1)
		req, e := http.Get(x + "/archive/refs/heads/" + sp[len(sp)-1] + ".zip")
		if e != nil {
			fmt.Println("Error retrieving content")
			return "", e
		}
		file, e := os.Create(sp[len(sp)-1])
		if e != nil {
			fmt.Println("Error creating file")
			return "", e
		}
		io.Copy(file, req.Body)
		return sp[len(sp)-1], nil
	}
	return "", nil
}
func Download(url string) {
	fmt.Println("Attempting to download file...")
	name, e := downloadFile_github(url)
	if e != nil {
		log.Fatal("Failed to download file/folder: ", e)
	}
	fmt.Println("Successfully downloaded repository zip file!")
	dir, e := os.Getwd()
	if e != nil {
		log.Fatal("Failed to find file/folder: ", e)
	}
	dir = dir + "/" + name
	zip, e := zip.OpenReader(dir)
	if e != nil {
		log.Fatal("Failed to unpack file/folder: ", e)
	}
	defer zip.Close()
	//arch, e := zip.Open(dir)
	//if e != nil {
	//	log.Fatal("Failed to unpack file/folder: ", e)
	//}
	for _, f := range zip.File {
		file, e := f.Open()
		if e != nil {
			log.Fatal("Failed to unpack file/folder: ", e)
		}
	}
	fmt.Println("Successfully unzipped files, Exit Code 0")
}
