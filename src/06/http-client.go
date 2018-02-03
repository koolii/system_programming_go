package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("[logger] send request to server")
	request.Write(conn)

	fmt.Println("[logger] get response from server")
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}

	fmt.Println("[logger] output response information")
	fmt.Println(string(dump))
}
