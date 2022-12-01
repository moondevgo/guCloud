package guCloud

// [Updating Google Sheets using Golang.](https://dev.to/mediocredevops/playing-with-google-sheets-api-using-golang-14en)
// [Sheets API Guides Go quickstart](https://developers.google.com/sheets/api/quickstart/go)

import (
	"fmt"
	"log"

	"google.golang.org/api/sheets/v4"
)

func SrvSheets(api_name, bot_nick, user_nick string) *sheets.Service {
	client := ApiClient(api_name, bot_nick, user_nick)
	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	return srv
}

func ReadSheet(srv *sheets.Service, spreadsheetId, readRange string) [][]interface{} {
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
		return nil
	} else {
		return resp.Values
	}
	return resp.Values
}

func WriteSheet(data [][]interface{}, srv *sheets.Service, spreadsheetId, writeRange string) {
	var vr sheets.ValueRange

	// myval := []interface{}{"One", "Two", "Three"}
	for _, d := range data {
		vr.Values = append(vr.Values, d)
	}

	_, err := srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
}
