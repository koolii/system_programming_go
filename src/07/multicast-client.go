package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listen tick server at 224.0.0.1:9999")
	// ListenMulticastUDPを使ってサーバを立てるがその時のアドレスの解決のために
	// net.ResolveUDPAddr()で変換を行わなければならない
	address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")

	if err != nil {
		panic(err)
	}

	listener, err := net.ListenMulticastUDP("udp", nil, address)
	defer listener.Close()

	buffer := make([]byte, 1500)

	for {
		// ReadFromUDP()はレスポンスで帰ってくるサーバのアドレスの型が
		// UDP専用のnet.UDPAddr型以外は普通のUDPクライアントで使ったReadFrom()とほぼ同じ
		length, remoteAddress, err := listener.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Server %v\n", remoteAddress)
		fmt.Printf("Now    %v\n", string(buffer[:length]))
	}
}
