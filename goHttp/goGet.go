package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "host:port")
		os.Exit(1)
	}

	url := os.Args[1]

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(3)
	}

	b, _ := httputil.DumpResponse(response, false)
	fmt.Println(string(b))

	contestTypes := response.Header["Contest-Type"]
	if !acceptableCharset(contestTypes) {
		fmt.Println("Cannot handle", contestTypes)
		os.Exit(4)
	}

	var buf [512]byte

	reader := response.Body

	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}
	os.Exit(0)
}

func acceptableCharset(contestTypes []string) bool {
	for _, cType := range contestTypes {
		if strings.Index(cType, "UTF-8") != -1 {
			return true
		}
	}
	return false
}
