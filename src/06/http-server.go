package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(
		strings.Join(
			request.Header["Accept-Encoding"],
			",",
		),
		"gzip",
	) != -1
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()

	// Accept後のソケットで何度も応答を返す為にループ
	// ここでループを追加することによって、TCPコネクションが張られた後も何度もリクエストを処理出来るようになる
	// => 無限ループ
	for {
		fmt.Println("[logger] From now on, it sets timeout into conn")
		// タイムアウトを設定(無限ループの終了条件)
		// ここで設定しておくことで、通信がしばらくないとタイムアウトのエラーでRead()の呼び出しを終了する
		// タイムアウトを設定しないとコネクションがブロックされたままになってしまう
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		// HTTPリクエストを分解
		// net.Connをbufio.Readerでラップして、それをhttp.ReadRequest()関数に渡している
		request, err := http.ReadRequest(bufio.NewReader(conn))

		if err != nil {
			// 標準のerrorインタフェースの上位互換であるnet.Errorインタフェースの構造体から取得可能
			// タイムアウト時のエラーはnet.Connが作成するが、それ以外のio.Readerは最初に発生したエラーをそのまま伝搬する
			// だからerrorからダウンキャストを行うことでタイムアウトかどうかを判断出来る
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				// タイムアウトもしくはソケットクローズ時は終了
				fmt.Println("[logger] Timeout")
				break
			} else if err == io.EOF {
				fmt.Println("[logger] EOF")
				break
			}
			// それ以外はエラーにする
			panic(err)
		}

		// デバッグ用の関数で、分解したHTTPリクエストを出力してくれる
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(dump))

		// HTTPレスポンスを生成
		// ContentLengthを返す為にcontent変数に返却する文字列で初期化する
		// Go言語だと、HTTP/1.0以下だったり、ContentLengthが無いとConnection: closeヘッダを付与する仕様がある
		// ここにあるようにヘッダの圧縮はされていない(HTTP/2で対応された)
		response := http.Response{
			StatusCode: 200,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
		}

		if isGZipAcceptable(request) {
			// gzip化
			content := "Hello World(gzipped)\n"
			var buffer bytes.Buffer
			// gzip.NewWriterで作成したio.Writerで圧縮を行う
			writer := gzip.NewWriter(&buffer)
			// エンコードした内容はbufferに書き込む
			io.WriteString(writer, content)
			writer.Close()

			response.Body = ioutil.NopCloser(&buffer)
			response.ContentLength = int64(buffer.Len())
			response.Header.Set("Content-Encoding", "gzip")
		} else {
			content := "Hello World\n"
			response.Body = ioutil.NopCloser(strings.NewReader(content))
			response.ContentLength = int64(len(content))
		}
		response.Write(conn)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("[logger] Server is running at localhost:8080")

	for {
		// コネクションを生成する時はここからスタートする
		fmt.Println("[logger] Create a new TCP connection and wait until a request come.")
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// リクエストの処理を一任
		go processSession(conn)

		// コネクションがタイムアウトすると、次のリクエストからは下記のログからスタートする
		fmt.Println("[logger] Finish this loop, but it hasn't handle response to a client yet")
	}
}
