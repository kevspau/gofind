package src

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"path/filepath"
)

var client = http.Client{}

func downloadFile_github(url string, wd string) (string, error) {
	sp := strings.Split(url, "/")
	primeBranch := true
	for _, v := range sp {
		if v == "tree" {
			primeBranch = true
		}
	}
	if primeBranch {
		req, e := http.Get("https://" + url + "/archive/refs/heads/master.zip")
		if e != nil {
			fmt.Println("Error retrieving content")
			return "", e
		}
		file, e := os.Create(filepath.Join(wd, sp[len(sp)-1] + "-master.zip"))
		if e != nil {
			fmt.Println("Error creating file")
			return "", e
		}
		io.Copy(file, req.Body)
		return filepath.Join(wd, sp[len(sp)-1] + "-master.zip"), nil
	} else if !primeBranch {
		x := strings.Replace("https://" + url, "/tree", "", -1)
		x = strings.Replace(x, "/"+sp[len(sp)-1], "", -1)
		req, e := http.Get(x + "/archive/refs/heads/" + sp[len(sp)-1] + ".zip")
		if e != nil {
			fmt.Println("Error retrieving content")
			return "", e
		}
		file, e := os.Create(filepath.Join(wd, sp[len(sp)-3] + "-" + sp[len(sp)-1] + ".zip")) //RUN RIGHT FUCKING NOW
		if e != nil {
			fmt.Println("Error creating file")
			return "", e
		}
		io.Copy(file, req.Body)
		return filepath.Join(wd, sp[len(sp)-3] + "-" + sp[len(sp)-1] + ".zip"), nil
	}
	return "", nil
}
func Download(url string) {
	root, e := os.Getwd()
	if e != nil {
		log.Fatal("Failed to find file/folder: ", e)
	}
	fmt.Println("Attempting to download file...")
	name, e := downloadFile_github(url, root)
	if e != nil {
		log.Fatal("Failed to download file/folder1: ", e)
	}
	fmt.Println("Successfully downloaded repository zip file!")
	fmt.Print("aaaaaaaaaa\n" + filepath.Join(root, name) + "\n\na\n")
	zip, e := zip.OpenReader(filepath.Join(root, name))
	if e != nil {
		log.Fatal("Failed to unpack file/folder2: ", e)
	}
	defer zip.Close()
	//arch, e := zip.Open(dir)
	//if e != nil {
	//	log.Fatal("Failed to unpack file/folder: ", e)
	//}
	for _, f := range zip.File {
		file, e := f.Open()
		defer file.Close()
		//if f.FileInfo().IsDir() && f.FileInfo().Name() == name {

		//}
		if e != nil {
			log.Fatal("Failed to unpack file/folder3: ", e)
		}
		if !(f.FileInfo().IsDir()) {
			nfile, e := os.Create(filepath.Join(root, f.FileInfo().Name()))
			if e != nil {
				log.Fatal("Failed to copy file/folder4: ", e)
			}
			io.Copy(nfile, file)
		} else {
			e := os.MkdirAll(filepath.Join(root, f.FileInfo().Name()), f.FileInfo().Mode())
			if e != nil {
				log.Fatal("Failed to copy file/folder5: ", e)
			}
		}
	}
	fmt.Println("Successfully unzipped files, Exit Code 0")
}
