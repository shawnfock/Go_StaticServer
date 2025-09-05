package main
import (
  "embed"
  "net/http"
)
//go:embed public
var content embed.FS
func main(){ http.Handle("/", http.FileServer(http.FS(content))); http.ListenAndServe(":8080", nil) }