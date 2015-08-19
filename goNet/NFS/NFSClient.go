package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	uiDir   = "dir"
	uiCd    = "cd"
	uiPwd   = "pwd"
	uiQuit  = "quit"
	uiLogin = "login"
)

const (
	DIR   = "DIR"
	CD    = "CD"
	PWD   = "PWD"
	LOGIN = "LOGIN"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "host")
		os.Exit(1)
	}

	host := os.Args[1]

	conn, err := net.Dial("tcp", host+":1201")
	checkError(err)

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimRight(line, "\t\r\n")
		if err != nil {
			break
		}

		strs := strings.SplitN(line, " ", 3)
		switch strs[0] {
		case uiLogin:
			loginRequest(conn, strs[1], strs[2])
		case uiDir:
			dirRequest(conn)
		case uiCd:
			if len(strs) != 2 {
				fmt.Println("cd <dir>")
				continue
			}
			fmt.Println("CD\"", strs[1], "\"")
			cdRequest(conn, strs[1])
		case uiPwd:
			pwdRequest(conn)
		case uiQuit:
			conn.Close()
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}

func dirRequest(conn net.Conn) {
	conn.Write([]byte(DIR + " "))
	var buf [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, _ := conn.Read(buf[0:])
		//s := string(buf[0:n])
		if string(buf[0:n]) == "NO" {
			fmt.Println("Please Login First")
			return
		}
		result.Write(buf[0:n])
		length := result.Len()
		contents := result.Bytes()
		if string(contents[length-4:]) == "\r\n\r\n" {
			fmt.Println(string(contents[0 : length-4]))
			return
		}
	}
}

func cdRequest(conn net.Conn, dir string) {
	conn.Write([]byte(CD + " " + dir))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	if s == "NO" {
		fmt.Println("Please Login First")
	} else if s != "OK" {
		fmt.Println("Failed to change dir")
	}
}

func pwdRequest(conn net.Conn) {
	conn.Write([]byte(PWD))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	if s == "NO" {
		fmt.Println("Please Login First")
	} else {
		fmt.Println("Current dir \"" + s + "\"")
	}
}

func loginRequest(conn net.Conn, name, pwd string) {
	conn.Write([]byte(LOGIN + " " + name + " " + pwd))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	if s == "Failed" {
		fmt.Println("Login Failed!")
	} else if s == "OK" {
		fmt.Println("Login Succeeded!")
	}
}
