package main

import (
  "fmt"
  "github.com/goji/zenazn/goji"
)

func main() {
  goji.Get("/", Root)
  goji.Post("/process_image", ProcessImage)
  goji.Serve()
}

func Root(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "Hello, and welcome to Picasso!")
}

func ProcessImage(w http.ResponseWriter, r *http.Request) {
  
}
