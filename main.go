package main

import (
	"encoding/base64"
	"fmt"
	"syscall/js"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func saveXlx(this js.Value, p []js.Value) interface{} {
	f := excelize.NewFile()
	// Create a new sheet.
	// Set value of a cell.
	tempo := 12
	tempo2 := 8
	f.SetCellValue("Sheet1", "A1", "Nome")
	f.SetCellValue("Sheet1", "B1", "Tempo")
	f.SetCellValue("Sheet1", "A2", "Teste Novo")
	f.SetCellValue("Sheet1", "A3", "desnehando")
	f.SetCellValue("Sheet1", "B2", tempo)
	f.SetCellValue("Sheet1", "B3", tempo2)

	err := f.AddChart("Sheet1", "C1", `{"type": "doughnut",
        "series": [{
            "name": "Sheet1!$A$1",
            "categories": "Sheet1!$A$2:$A$3",
            "values": "Sheet1!$B$2:$B$3"
        }]}`)

	if err != nil {
		fmt.Println(err)
	}

	buffer, err := f.WriteToBuffer()

	bs := make([]byte, 0)
	bs = append(bs, 0, 1, 255, 140, 200)

	encodedStr := base64.StdEncoding.EncodeToString(buffer.Bytes())
	return encodedStr
}

func shareBytes(this js.Value, p []js.Value) interface{} {
	dados := make([]byte, 0)
	dados = append(dados, 25, 50, 130, 85, 20)

	outArray := bytesToJSArray(dados)

	return outArray
}

func bytesToJSArray(input []byte) interface{} {
	arrayConstructor := js.Global().Get("Uint8Array")
	outArray := arrayConstructor.New(len(input))
	js.CopyBytesToJS(outArray, input)
	return outArray
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("xlx", js.FuncOf(saveXlx))
	js.Global().Set("shareBytes", js.FuncOf(shareBytes))
	<-c
}
