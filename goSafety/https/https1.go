package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	SERVER_PORT       = 5050
	SERVER_DOMAIN     = "localhost"
	RESPONSE_TEMPLSTE = "hello"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLSTE)))
	w.Write([]byte(RESPONSE_TEMPLSTE))
}

func main() {
	http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMAIN, SERVER_PORT), rootHandler)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "rui.crt", "rui.key", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS:", err.Error())
	}
}
