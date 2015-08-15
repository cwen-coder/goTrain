package main

import (
	"encoding/asn1"
	"fmt"
	"os"
	"time"
)

func main() {
	mdata, err := asn1.Marshal(13)
	checkError(err)
	var n int
	_, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)
	fmt.Println("After marshal/unmarshal:", n)

	t := time.Now()
	m, err := asn1.Marshal(t)
	var newtime = new(time.Time)
	_, err = asn1.Unmarshal(m, newtime)
	checkError(err)
	fmt.Println(newtime)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error :%s\n", err.Error())
		os.Exit(1)
	}
}
