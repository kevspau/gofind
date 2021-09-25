package main

import (
	"os"
	"strings"
	"github.com/csharpdf/gofind/src"
	"fmt"
)

func main() {
	if os.Args[1] == "find" {
		if strings.HasPrefix(os.Args[2], "github.com") || strings.HasPrefix(os.Args[2], "https://github.com") {
			src.Download(os.Args[1])
		} else {
			fmt.Println("Unsupported domain given. Current supported domains are...\n\n [https://]github.com")
		}
	} else if os.Args[1] == "help" {
		src.Help()
	} else if os.Args[1] == "package" {
		if os.Args[2] == "" {
			src.PackageHelp()
		} else if os.Args[2] == "init" {
			//make .find file parser first
		} else if os.Args[2] == "find" {
			//above
		} else if os.Args[2] == "search" {
			//above and connect database(sql or mongodb) or make .find files automate builds
		}
	}
}
