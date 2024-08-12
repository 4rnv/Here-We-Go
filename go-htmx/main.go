package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Film struct {
	Name     string
	Director string
	Year     string
	Rating   float32
}

func readjson(filename string) ([]Film, error) {
	byteValue, err := os.ReadFile(filename)
	var films []Film
	if err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal(byteValue, &films); err != nil {
		fmt.Println(err)
	}
	return films, nil
}

func ichi(w http.ResponseWriter, r *http.Request) {
	films, _ := readjson("films.json")
	//fmt.Println(films)
	plate := template.Must(template.ParseFiles("index.html"))
	plate.Execute(w, films)
}

func ni(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HTMX Post was triggered.", r.Header.Get("HX-Request"))
	query := r.PostFormValue("search-query")
	fmt.Println(query)

	films, _ := readjson("films.json")
	var results []Film
	for _, film := range films {
		if (strings.Contains(strings.ToLower(film.Name), strings.ToLower(query))) || (strings.Contains(strings.ToLower(film.Director), strings.ToLower(query))) {
			results = append(results, film)
		}
	}
	fmt.Println(results)
	if err := renderResults(w, results); err != nil {
		fmt.Println(err)
	}
}

func renderResults(w http.ResponseWriter, films []Film) error {
	tmpl := `    
	{{ range . }}
    <div class="film">
    <h2>{{.Name}}</h2>
    <p>Director: {{.Director}}</p>
    <p>Year: {{.Year}}</p>
    <p>Rating: {{.Rating}}</p>
    </div>
    {{ end }}`

	t, err := template.New("results").Parse(tmpl)
	if err != nil {
		return err
	}

	return t.Execute(w, films)
}

func main() {
	fmt.Println("Sneed")
	http.HandleFunc("/", ichi)
	http.HandleFunc("/search/", ni)
	port := ":10000"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error running server", err)
		log.Fatal(err)
	} else {
		fmt.Println("Server running on port ", port)
	}
}
