package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/thecsw/haruhi"
)

type StructJSON struct {
	Id        string `json:"_id"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
	Show      string `json:"show"`
}

func main() {
	requrl := "https://yurippe.vercel.app/api/quotes"
	params := url.Values{
		"show":   {"haruhi"},
		"random": {"1"},
	}

	//String
	response, err :=
		haruhi.
			URL(requrl).
			Params(params).
			Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)

	//JSON
	var responseJSON []StructJSON

	errr :=
		haruhi.
			URL(requrl).
			Params(params).
			ResponseJson(&responseJSON)
	if errr != nil {
		fmt.Println(errr)
	}

	jsonBytes, errrr := json.MarshalIndent(responseJSON, "", "  ")
	if errrr != nil {
		fmt.Println(errrr)
		return
	}
	fmt.Println(string(jsonBytes))
}
