package main

import (
	"log"
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir("."))
	err := http.ListenAndServeTLS(":8001", "rui.crt", "rui.key", h)
	if err != nil {
		log.Fatal(err.Error())
	}
}
