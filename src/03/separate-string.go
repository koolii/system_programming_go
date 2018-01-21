package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `1行目
2行目
3行目`

func normal() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)

		if err == io.EOF {
			break
		}
	}
}

// normalのようにEOFを気にせず書くなら、bufio.Scannerは便利
// ただ、出力すれば分かるが、分割文字（今回は`\n`）がなくなっている
// ※ bufio.Scannerは分割関数を指定することができる（例えば、単語区切りの`scanner.Split(busio.ScanWords`）を設定することで実現可能)
func simple() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}

func main() {
	fmt.Println("※ Take care of EOF character")
	normal()
	fmt.Println("※ Don't take care of EOF character")
	simple()
}
