package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

func main() {
	eightBitData := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	bb := &bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, bb)
	encoder.Write(eightBitData)
	encoder.Close()
	fmt.Println(bb)
	decoder := base64.NewDecoder(base64.StdEncoding, bb)
	buf := make([]byte, 1)
	for {
		n, err := decoder.Read(buf)
		if err != nil {
			break
		}
		if n == 0 {
			break
		}
		fmt.Print(buf)
	}

}
