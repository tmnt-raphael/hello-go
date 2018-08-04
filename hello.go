package main

import (
  "fmt"
  "net/http"
  "github.com/tmnt-raphael/kenwebutil"
)

func main() {
  fmt.Printf("Go to localhost:8080/some-path.\n")
  fmt.Printf("The path of the URL will be returned.\n")

  http.HandleFunc("/", kenwebutil.EchoPath)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}