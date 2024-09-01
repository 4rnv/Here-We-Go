package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var endpoint = "https://api.cognitive.microsofttranslator.com/translate?api-version=3.0"
var apikey = "<YOUR API KEY GOES HERE>"
var location = "southeastasia"

func translate(text string, lang string) any {
	if text == "" || lang == "" {
		return errors.New("Either language or text not specified")
	}
	u, _ := url.Parse(endpoint)
	q := u.Query()
	q.Add("to", lang)
	u.RawQuery = q.Encode()
	body := []struct {
		Text string
	}{
		{Text: text},
	}
	b, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Ocp-Apim-Subscription-Key", apikey)
	req.Header.Add("Ocp-Apim-Subscription-Region", location)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//Decode the JSON response
	var result interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	// Format and print the response to terminal
	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("%s\n", prettyJSON)
	return prettyJSON
}

func main() {
	var text, lang string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the text you want to translate")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Println("Enter the language code for the language you want to translate into, for example ja(Japanese) or de(German)")
	lang, _ = reader.ReadString('\n')
	lang = strings.TrimSpace(lang)
	if lang != "" && text != "" {
		translate(text, lang)
	}
}
