package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "Http://host:port/page")
		os.Exit(1)
	}

	url, err := url.Parse(os.Args[1])
	checkError(err)
	client := &http.Client{}
	request, err := http.NewRequest("Get", url.String(), nil)
	request.Header.Add("Accept-Charset", "UTF-8;q=1,ISO-8859-1;q=0")
	checkError(err)
	response, err := client.Do(request)
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	chSet := getCharset(response)
	fmt.Println("got charset %s\n", chSet)
	if chSet != "UTF-8" {
		fmt.Println("Cannot handle", chSet)
		os.Exit(4)
	}

	var buf [512]byte
	reader := response.Body
	fmt.Println("go body")
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Println(buf[0:n])
	}
	os.Exit(0)

}

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		return "UTF-8"
	}

	idx := strings.Index(contentType, "charset")
	if idx == -1 {
		return "UTF-8"
	}
	return strings.Trim(contentType[idx:], "")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}
