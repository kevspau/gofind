package main

import "fmt"

func help() {
	fmt.Println("Commands:\n\nhelp - Lists current commands\nfind [github repo url] - downloads the specified github url into the working directory, and then unpacks it\npackage - Lists current commands that provide management for gofind packages.")
}

func packageHelp() {
	fmt.Println("The gofind package system uses custom .find files to manage packages. Once given a proper .find file, the new package will be added to the gofind database for others to download. Make sure to include all wanted files within the same directory of the .find file, or they will be excluded.\n\nCurrent commands:\n\npackage init - initializes a gofind package, creating a .find file in the process.\npackage find [package name] - downloads the specified package in the named directory\npackage search [package name] - Searches the gofind database for the specified database, and returns a boolean statement to the command line.")
}