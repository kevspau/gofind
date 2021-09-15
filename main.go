package main

import (
  "fmt"
  "os"
  "io"
  "net/http"
  "log"
  "strings"
)
var client = http.Client{}
func downloadFile(url string) *error {
  sp := strings.Split(url, "/")
  err := os.Mkdir(sp[len(sp) - 1], 1)
  if err != nil {
    return &err
  }
  os.Chdir("./" + sp[len(sp) - 1])
  var newg string
  if strings.HasPrefix(url, "github.com") || strings.HasPrefix(url, "https://github.com") {
    newg = url + "/archives/reps/heads/master.zip"
  } else {
    log.Fatal("Unsupported domain given.")
  }
  file, err := os.Create(sp[len(sp) - 1] + ".zip")
  os.Chdir("../")
  if err != nil {
    return &err
  }
  g, err := client.Get(newg)
  if err != nil {
    return &err
  }
  io.Copy(file, g.Body)
  return nil
}
func main() {
  fmt.Println("Attempting to download file...")
   err := downloadFile(os.Args[1])
   if err != nil {
     log.Fatal("Failed to download file/folder:", err)
   }
}
