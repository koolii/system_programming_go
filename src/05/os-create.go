package main

import "os"

func main() {
	file, err := os.Create("os-create.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// os.Create => os.OpenFile() => syscall.Open()
	// 最終的に Syscall()関数を呼び出して、OSに実行してもらうシステムコールを決定する
	file.Write([]byte("system call example"))
}
