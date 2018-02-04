package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println("[logger] Make a request to server")
	request, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("[logger] Send a request to server")
	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	fmt.Println("[logger] Get response from server")
	reader := bufio.NewReader(conn)
	response, err := http.ReadResponse(reader, request)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(response, false)
	if err != nil {
		panic(err)
	}

	fmt.Println("[logger] Output response information")
	fmt.Println(string(dump))

	if len(response.TransferEncoding) < 1 || response.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}

	for {
		sizeStr, err := reader.ReadBytes('\n')
		if err != io.EOF {
			break
		}

		// 16進数のサイズをパース
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if size == 0 {
			break
		}

		if err != nil {
			panic(err)
		}

		line := make([]byte, int(size))
		reader.Read(line)
		reader.Discard(2)
		fmt.Printf("%d bytes: %s\n", size, string(line))
	}
}
