package main

import (
  "flag"
  "os"
  "io"
  "net/http"
  "log"
  "strings"
)
var client = http.Client{}
func downloadFile(url string) {
  g, err := client.Get(url)
  sp := strings.Split(url, "/")
  if err != nil {
    log.Fatal(err)
  }
  err = os.Mkdir(sp[len(sp)], 0777)
  if err != nil {
    log.Fatal(err)
  }
  os.Chdir("./" + sp[len(sp)])
  if strings.HasPrefix(url, "github.com") || strings.HasPrefix(url, "https://github.com") {
    g, err = client.Get(url + "/archives/reps/heads/master.zip")
  } else {
    log.Fatal("Unsupported domain given. Current supported domains include...\n[https://]github.com")
  }
  file, err := os.Create(sp[len(sp)])
  os.Chdir("../")
  if err != nil {
    log.Fatal(err)
  }
  io.Copy(file, g.Body)
}
func main() {
  downloadFile(flag.Args()[0])
}
