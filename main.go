package main

import (
	"fmt"
	"net/http"
	"time"
)

func test(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>Root path</h1>")
	case "/api":
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"sneed\": \"feed\", \"chuck\": \"feeduck\"}")
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error 404")
	}
	fmt.Println(r.Method, r.URL)
}

func customServerConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Resolving le timeout")
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, "Custom server config route")
	fmt.Fprintf(w, "This is the server page, timeout This should not be visible")
}

func main() {
	http.HandleFunc("/", test)
	http.HandleFunc("/server", customServerConfig)
	server := http.Server{
		Addr:         ":8000",
		Handler:      nil,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	serverStatus := server.ListenAndServe()
	if serverStatus != nil {
		fmt.Println("Error running le server: ", serverStatus)
	}
}
