package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed index.html
var indexHTML embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(indexHTML)))

	addr := ":8000"
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
