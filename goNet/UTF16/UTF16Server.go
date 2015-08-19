package main

import (
	"fmt"
	"net"
	"os"
	"unicode/utf16"
)

const BOM = '\ufffe'

func main() {
	service := ":1210"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		str := "j'ai arrêté"

		shorts := utf16.Encode([]rune(str))
		writeShorts(conn, shorts)
		conn.Close()
	}
}

func writeShorts(conn net.Conn, shorts []uint16) {
	var bytes [2]byte
	bytes[0] = BOM >> 8
	bytes[1] = BOM & 255
	_, err := conn.Write(bytes[0:])
	if err != nil {
		return
	}

	for _, v := range shorts {
		bytes[0] = byte(v >> 8)
		bytes[1] = byte(v & 255)
		_, err = conn.Write(bytes[0:])
		if err != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}
