package main

import (
	"fmt"
	"io"
	"os"
)

// go run src/03/standard-input.go < src/03/standard-input.go
func main() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)

		if err == io.EOF {
			fmt.Println("EOF")
			break
		}

		fmt.Printf("size=%d, input='%s'\n", size, string(buffer))
	}
}
