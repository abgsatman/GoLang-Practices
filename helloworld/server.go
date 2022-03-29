package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type Lego struct {
	Type string
}

func server() {

	fmt.Println("server.go file is running............")

	http.HandleFunc("/", indexHandler)      //index
	http.HandleFunc("/hc", hcHandler)       //health check controller
	http.HandleFunc("/json", jsonHandler)   //json responder
	http.HandleFunc("/print", printHandler) //html writer

	http.ListenAndServe(":18080", nil) //the server is always on by this way
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	value, error := fmt.Fprintf(w, "Hello, World!\n")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if error == nil {
		fmt.Fprintf(w, "%d bytes written\n", value)
	} else {
		fmt.Fprintf(w, "An error occured.")
	}
}

func hcHandler(w http.ResponseWriter, r *http.Request) {
	response := w
	if w == nil {
		fmt.Fprintf(response, "hc is not working!")
	} else {
		fmt.Fprintf(response, "hc is working...")
	}
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	lego := Lego{Type: "Lego Creator"}
	json.NewEncoder(w).Encode(lego)
}

func printHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("src/lego.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load the html file")
	}

	lego := Lego{Type: "Lego Creator"}
	t.Execute(w, lego)
}
