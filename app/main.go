package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type RequestLine struct {
	Method  string
	Target  string
	Version string
}

type Headers struct {
	Host   string
	Agent  string
	Accept string
}

const (
	Ok       = "HTTP/1.1 200 OK\r\n\r\n"
	NotFound = "HTTP/1.1 404 Not Found\r\n\r\n"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// TODO: Uncomment the code below to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	CRLF := "\r\n"
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Error reading connection buf: %v", err)
			continue
		}

		reqStr := string(buf)
		reqStrSplit := strings.Split(reqStr, CRLF)

		requestLine := parseRequestLine(reqStrSplit[0])
		headers := parseHeaders(reqStrSplit[1])

		fmt.Println(requestLine, headers)

		if requestLine.Target == "/" {
			conn.Write([]byte(Ok))
		} else {
			conn.Write([]byte(NotFound))
		}

		// response := "HTTP/1.1 200 OK\r\n\r\n"
		// conn.Write([]byte(response))
	}
}

func parseRequestLine(r string) *RequestLine {
	rSplit := strings.Split(r, " ")
	return &RequestLine{
		Method:  rSplit[0],
		Target:  rSplit[1],
		Version: rSplit[2],
	}
}

func parseHeaders(h string) *Headers {
	hSplit := strings.Split(h, " ")
	return &Headers{
		Host: hSplit[0],
		// Agent:  hSplit[1],
		// Accept: hSplit[2],
	}
}
