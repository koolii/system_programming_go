package main

import "fmt"

func main() {
	// byteArrayは {0x41, 0x53, 0x43, 0x49, 0x49}
	byteArray := []byte("ASCII")

	// strは "ASCII"
	str := string([]byte{0x41, 0x53, 0x43, 0x49, 0x49})

	fmt.Println(byteArray)
	fmt.Println(str)
}
