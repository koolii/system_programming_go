package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// big-endian data
	data := []byte{0x0, 0x0, 0x27, 0x10}

	// convert
	var i int32
	// 現在のCPUに合わせて変換する
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Println(i)
}
