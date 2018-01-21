package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// ちょっと中身の実装がはっきり理解できてない
// チャンクを表示する
func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes) \n", string(buffer), length)
}

// PNGファイルをチャンク毎にバイト配列を生成する
func readChunks(file *os.File) []io.Reader {
	// 格納する配列を宣言
	var chunks []io.Reader

	// PNGの最初のシグネチャ分を飛ばす
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		// ここでBigEndian形式で読み込む、lengthには、データ長さが挿入される
		// おそらく固定で、32バイトとかの固定バイトのバッファを作成して、読み込みをしている気がする
		err := binary.Read(file, binary.BigEndian, &length)
		if err != nil {
			break
		}

		// なんでlengthに`+12`するのかわからない(おそらくPNGのデータチャンクの仕様だとは思う)
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))

		// 次のチャンクの先頭に移動、チャンク名(4バイト)＋データ長＋CRC(4バイト)先に移動 => 8バイト
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func main() {
	file, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
