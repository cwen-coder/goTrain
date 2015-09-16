package main

import (
	"net/http"
)

func main() {
	myHandler := http.HandleFunc(func(rw http.ResponseWriter, request *http.Request) {
		rw.WriteHeader(http.StatusNoContent)
	})

	http.ListenAndServe(":6666", myHandler)
}
