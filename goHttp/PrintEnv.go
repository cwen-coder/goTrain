package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServer(http.Dir("/home/yin_cwen"))
	http.Handle("/", fileServer)
	http.HandleFunc("cgi-bin/printenv", printEnv)
	err := http.ListenAndServe(":7777", nil)
	checkError(err)
}

func printEnv(writer http.ResponseWriter, req *http.Request) {
	env := os.Environ()
	writer.Write([]byte("<h1>Environment</h1>\n<pre>"))
	for _, v := range env {
		writer.Write([]byte(v + "\n"))
	}
	writer.Write([]byte("</pre>"))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fail error", err.Error())
		os.Exit(1)
	}
}
