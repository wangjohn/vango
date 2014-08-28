package main

import (
  "fmt"
  "io"
  "net/http"
  "regexp"
  "strconv"
  "image"
  _ "image/png"
  _ "image/jpeg"
  _ "image/gif"

  "github.com/zenazn/goji"
  "github.com/wangjohn/vango/primary_color"
)

const (
  maxContentLength int = 5140000 // 5MB
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
  contentTypes := r.Header["Content-Type"]
  contentTypeRegex, _ := regexp.Compile("multipart/form-data")

  for _, contentType := range contentTypes {
    if !contentTypeRegex.MatchString(contentType) {
      message := fmt.Sprintf("Bad content type: `%s` is not currently not supported", contentType)
      http.Error(w, message, http.StatusBadRequest)
    }
  }

  contentLength, _ := strconv.Atoi(r.Header["Content-Length"][0])
  if contentLength > maxContentLength {
    message := fmt.Sprintf("Bad content length: `%d` exceeds the maximum file size", contentLength)
    http.Error(w, message, http.StatusBadRequest)
  }

  err := r.ParseMultipartForm(int64(maxContentLength))
  if err != nil {
    message := fmt.Sprintf("Error parsing multipart-form: %s", err.Error())
    http.Error(w, message, http.StatusBadRequest)
  }

  imgFiles, imgFileExists := r.MultipartForm.File["file"]
  if !imgFileExists {
    message := "Error parsing multipart-form: must specify a `file` field"
    http.Error(w, message, http.StatusBadRequest)
  }

  if len(imgFiles) > 1 {
    message := "Only expected a single file to be uploaded, not multiple"
    http.Error(w, message, http.StatusBadRequest)
  }

  multipartImgFile, err := imgFiles[0].Open()
  if err != nil {
    message := fmt.Sprintf("Error opening file: %s", err.Error())
    http.Error(w, message, http.StatusBadRequest)
  }

  img, _, err := image.Decode(multipartImgFile)
  if err != nil {
    message := fmt.Sprintf("Error decoding image file: %s", err.Error())
    http.Error(w, message, http.StatusBadRequest)
  }

  handleImage(img)
}

func handleImage(img image.Image) {
  primaryColor.PrimaryColor(img)
}
