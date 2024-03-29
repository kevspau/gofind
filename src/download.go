package src

import ( 
	"strings"
	"fmt"
	"os"
	"net/http"
	"path/filepath"
	"log"
	"archive/zip"
	"io"
)

var client = http.Client{}

func downloadFile(url string) error {
	sp := strings.Split(url, "/")
	var newg string
	newg = url + "/archive/refs/heads/master.zip"
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
func Download(url string) {
	inDir := false
	fmt.Println("Attempting to download file...")
	err := downloadFile(url)
	if err != nil {
		log.Fatal("Failed to download file/folder: ", err)
	}
	fmt.Println("Successfully downloaded repository zip file!")
	link := strings.Split(url, "/")
	zip, err := zip.OpenReader(filepath.Join(link[len(link)-1], "/main.zip"))
	if err != nil {
		log.Fatal("Failed to unpack file/folder: ", err)
	}
	defer zip.Close()
	for _, fi := range zip.File {
		fmt.Printf("Unpacking %s..", fi.Name)
		if !fi.FileHeader.FileInfo().IsDir() {
			if !inDir {
				file, err := fi.Open()
				if err != nil {
					log.Fatal("Failed to unpack file/folder: ", err)
				}
				Foo, err := os.Create(filepath.Join(link[len(link)-1], fi.Name))
				if err != nil {
					log.Fatal("Failed to unpack file/folder: ", err)
				}
				readerr, err := fi.Open()
				if err != nil {
					log.Fatal("Failed to unpack file/folder: ", err)
				}
				io.Copy(Foo, readerr)
				fmt.Printf("(Size %d) [Done]\n", fi.FileInfo().Size())
				Foo.Close()
			} else {
				os.Create(filepath.Join(filepath.Join(link[len(link)-1], dir), f))
				fil, _ := os.Open(f)
				io.Copy(fil, file)
				fmt.Print(" [Done]\n")
				fil.Close()
				inDir = false
			}
		} else {
			dir, _ := filepath.Split(fi.Name)
			os.MkdirAll(filepath.Join(link[len(link)-1], dir), 1)
			inDir = true
		}
	}
	fmt.Println("Successfully unzipped files, Exit Code 0")
}