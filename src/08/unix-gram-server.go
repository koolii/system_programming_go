package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	pat := filpath.Join(os.TempDir(), "unixdomainsocket-sample")
	os.Remove(path)
	conn, err := net.ListenPacket("unixgram", path)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	buffer := make([]byte, 1500)

	fmt.Println("Server is running at " + path)

	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Received from %v: %v\n", remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)

		if err != nil {
			panic(err)
		}
	}
}
