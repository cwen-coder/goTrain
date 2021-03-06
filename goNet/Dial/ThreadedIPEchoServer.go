package main

import (
	"fmt"
	"net"
	"os"
	//"time"
)

func main() {
	service := ":1202"
	listener, err := net.Listen("tcp", service)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//daytime := time.Now().String()
		//conn.Write([]byte(daytime)) // don't care about return value
		go handleClient(conn)
		//conn.Close() // we're finished with this client
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	//daytime := time.Now().String()
	//conn.Write([]byte(daytime))
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:]))
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
	//time.Sleep(100 * time.Millisecond)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
