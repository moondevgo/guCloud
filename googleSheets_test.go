package guCloud

import (
	"fmt"
	"testing"
)

// import (
// 	"github.com/moondevgo/go-mods/cloud"
// )

// var srv *sheets.Service
// srv = SrvSheets("sheets", "moonsats", "")

func TestSrvSheets(t *testing.T) {
	srv := SrvSheets("sheets", "moonsats", "")

	if srv == nil {
		t.Error("Wrong result")
	}
}

func TestReadSheet(t *testing.T) {
	srv := SrvSheets("sheets", "moonsats", "")
	spreadsheetId := "1TofihRh87iKRsNOCrjQusCgStPkIDx-wOqZa1xee-SI"
	// readRange := "sheet2!A1:C17"
	readRange := "candles"
	res := ReadSheet(srv, spreadsheetId, readRange)

	if res == nil {
		t.Error("Wrong result")
	}
	fmt.Printf("%v", res)
}

func TestWriteSheet(t *testing.T) {
	srv := SrvSheets("sheets", "moonsats", "")
	spreadsheetId := "1TofihRh87iKRsNOCrjQusCgStPkIDx-wOqZa1xee-SI"
	writeRange := "sheet3!A1"
	data := [][]interface{}{
		{"One", "Two", "Three"},
		{1, 2, 3},
	}
	WriteSheet(data, srv, spreadsheetId, writeRange)
}
