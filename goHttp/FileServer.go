package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServer(http.Dir("/home/yin_cwen"))
	err := http.ListenAndServe(":7777", fileServer)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fail error", err.Error())
		os.Exit(1)
	}
}
