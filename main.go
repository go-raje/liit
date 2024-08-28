package main

import (
	"fmt"
	"os"
	"io"
	"net/http"
)

func downloadFile(filepath string, url string) (err error) {

  // Create the file
  out, err := os.Create(filepath)
  if err != nil  {
    return err
  }
  defer out.Close()

  // Get the data
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  // Writer the body to file
  _, err = io.Copy(out, resp.Body)
  if err != nil  {
    return err
  }

  return nil
}

func main() {
	err := downloadFile(os.Args[1], "https://pkgs.merelinux.org/core/os/x86_64/acl-2.3.2-1-x86_64.pkg.tar.xz")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
