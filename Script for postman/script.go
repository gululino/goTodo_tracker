package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/xuri/excelize/v2"
)

type PostmanCollection struct {
	Info struct {
		Name   string `json:"name"`
		Schema string `json:"schema"`
	} `json:"info"`
	Item []PostmanItem `json:"item"`
}

type PostmanItem struct {
	Name    string         `json:"name"`
	Request PostmanRequest `json:"request"`
}

type PostmanRequest struct {
	Method string          `json:"method"`
	Header []PostmanHeader `json:"header"`
	Body   PostmanBody     `json:"body"`
	URL    PostmanURL      `json:"url"`
}

type PostmanHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PostmanBody struct {
	Mode string `json:"mode"`
	Raw  string `json:"raw"`
}

type PostmanURL struct {
	Raw  string   `json:"raw"`
	Host []string `json:"host"`
	Path []string `json:"path"`
}

type RequestBody struct {
	TransactionDescription string `json:"transaction_description"`
	TransactionType        string `json:"transaction_type"`
	CardNumber             string `json:"card_number"`
	CardExpiration         string `json:"card_expiration"`
	CardVerification       string `json:"card_verification"`
	StreetAddress1         string `json:"street_address1"`
	Zip                    string `json:"zip"`
	TransactionAmount      string `json:"transaction_amount"`
	TenderType             string `json:"tender_type"`
	FinalAuthorization     int    `json:"final_authorization"`
	AccountID              string `json:"account_id"`
	APIAccessKey           string `json:"api_accesskey"`
	ResponseFormat         string `json:"response_format"`
}

var (
	url     = "https://tul-pc-dev-mjoseph-21.cardconex.local/api/qsapi/3.8"
	headers = []PostmanHeader{
		{"content-type", "application/json"},
		{"Cookie", "visid_incap_160084=OK3WHDo6R6mql8ohx90bVl2FnGgAAAAAQUIPAAAAAABU07PbvRBbnZ6fmiWBlhJx; visid_incap_168247=bv3+G+0XR0OFCvnukXSZ9+ZfvGcAAAAAQUIPAAAAAAD1vIz3x6+ww1KX9LAUz4B6"},
	}
	constants = map[string]interface{}{
		"card_expiration":     "1230",
		"api_accesskey":       "1234566789012",
		"account_id":          "184572907838",
		"response_format":     "JSON",
		"tender_type":         "CARD",
		"final_authorization": 1,
	}
)

// Extract helpers
func extractCardNumber(s string) string {
	re := regexp.MustCompile(`PAN\s*=\s*(\d+)`)
	match := re.FindStringSubmatch(s)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func extractAmount(s string) string {
	re := regexp.MustCompile(`amount of[^\d]*(\d+(\.\d{1,2})?)`)
	match := re.FindStringSubmatch(s)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func extractZip(s string) string {
	re := regexp.MustCompile(`\b(\d{5})\b`)
	match := re.FindStringSubmatch(s)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func extractAddress(s string) string {
	zip := extractZip(s)
	if zip == "" {
		return ""
	}
	idx := strings.Index(s, zip)
	if idx > 0 {
		beforeZip := s[:idx]
		parts := strings.Split(beforeZip, "\n")
		return strings.TrimSpace(parts[len(parts)-1])
	}
	return ""
}

func extractTransactionType(name string) string {
	re := regexp.MustCompile(`\d+\s+([A-Za-z]+)`)
	match := re.FindStringSubmatch(name)
	if len(match) > 1 {
		return strings.ToUpper(match[1])
	}
	return "SALE"
}

func main() {
	f, err := excelize.OpenFile("Test_Script_2335_2025_09_25_15_20_MMotto.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Fatal(err)
	}

	// header index mapping
	headersRow := rows[0]
	colIndex := make(map[string]int)
	for i, h := range headersRow {
		colIndex[h] = i
	}

	collection := PostmanCollection{}
	collection.Info.Name = "MOTO Transactions MSR 2025"
	collection.Info.Schema = "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"

	for _, row := range rows[1:] {
		if len(row) < len(headersRow) {
			continue
		}

		testCaseName := row[colIndex["Test Case Name"]]
		desc := row[colIndex["Transaction Description"]]
		triggers := row[colIndex["Transaction Triggers"]]
		instructions := row[colIndex["Instructions"]]

		body := RequestBody{
			TransactionDescription: desc,
			TransactionType:        extractTransactionType(testCaseName),
			CardNumber:             extractCardNumber(triggers),
			CardExpiration:         constants["card_expiration"].(string),
			CardVerification:       "1234",
			StreetAddress1:         extractAddress(instructions),
			Zip:                    extractZip(instructions),
			TransactionAmount:      extractAmount(instructions),
			TenderType:             constants["tender_type"].(string),
			FinalAuthorization:     constants["final_authorization"].(int),
			AccountID:              constants["account_id"].(string),
			APIAccessKey:           constants["api_accesskey"].(string),
			ResponseFormat:         constants["response_format"].(string),
		}

		rawBody, _ := json.MarshalIndent(body, "", "  ")

		item := PostmanItem{
			Name: testCaseName,
			Request: PostmanRequest{
				Method: "POST",
				Header: headers,
				Body: PostmanBody{
					Mode: "raw",
					Raw:  string(rawBody),
				},
				URL: PostmanURL{
					Raw:  url,
					Host: []string{"cert", "payconex", "net"},
					Path: []string{"api", "qsapi", "3.8"},
				},
			},
		}
		collection.Item = append(collection.Item, item)
	}

	out, err := os.Create("Moto_postman_collection.json")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")
	if err := enc.Encode(collection); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Postman collection generated: Moto_postman_collection.json")
}
