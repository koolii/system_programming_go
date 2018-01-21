package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var csvSource = `07543,"97906","9790622","ﾌｸｼﾏｹﾝ","ﾌﾀﾊﾞｸﾞﾝﾄﾐｵｶﾏﾁ","ｹｶﾞﾔ(ﾏｴｶﾜﾊﾗ232-244､311､312､337-862ﾊﾞﾝﾁ","福島県","双葉郡富岡町","毛萱（前川原２３２〜２４４、３１１、３１２、３３７〜８６２番地",1,1,0,0,0,0
07543,"97906","9790622","ﾌｸｼﾏｹﾝ","ﾌﾀﾊﾞｸﾞﾝﾄﾐｵｶﾏﾁ","ﾄｳｷｮｳﾃﾞﾝﾘｮｸﾌｸｼﾏﾀﾞｲ2ｹﾞﾝｼﾘｮｸﾊﾂﾃﾞﾝｼｮｺｳﾅｲ)","福島県","双葉郡富岡町","〔東京電力福島第二原子力発電所構内〕）",1,1,0,0,0,0`

func main() {
	// csv.NewReaderでio.Readerを返す
	reader := strings.NewReader(csvSource)
	csvReader := csv.NewReader(reader)

	for {
		// 行の情報（文字列の配列）を返す
		// ReadAll()でそれが更に配列になったものを一度に返すこともできる
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(line[2], line[6:9])
	}
}
