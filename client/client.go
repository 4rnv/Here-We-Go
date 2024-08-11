package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
)

type User struct {
	Username string
	Password string
}

func client() {
	response, erro := http.Get("https://obscure-capybara-q7qj9769w4rph567-1488.app.github.dev")
	if erro != nil {
		fmt.Println(erro)
	} else {
		userdata, status := io.ReadAll(response.Body)
		fmt.Println(string(userdata))
		fmt.Println(status)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	var fileName = "index.html"
	fmt.Println(r)
	template, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error parsing file", err)
		return
	}
	err = template.ExecuteTemplate(w, fileName, "Login")
	if err != nil {
		fmt.Println("Error when templating", err)
		return
	}
}

func LoginSubmit(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(w)
	var SessionUser User
	SessionUser.Username = strings.TrimSpace(r.FormValue("username"))
	SessionUser.Password = r.FormValue("password")
	fmt.Println(SessionUser)
	fmt.Fprintf(w, "Hello %s", SessionUser.Username)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		login(w, r)
	case "/login_submit":
		LoginSubmit(w, r)
	default:
		http.NotFound(w, r)
	}
	fmt.Println("Server running")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}
