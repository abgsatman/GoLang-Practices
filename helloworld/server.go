package main

import (
	"fmt"
	"net/http"
)

func server() {

	fmt.Println("server.go file is running............")

	index()
	hc()

	http.ListenAndServe(":80", nil)
}

func index() {
	//index endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		value, error := fmt.Fprintf(w, "Hello, World!\n")
		if error == nil {
			fmt.Fprintf(w, "%d bytes written\n", value)
		} else {
			fmt.Fprintf(w, "An error occured.")
		}
	})
}

func hc() {
	//health check
	http.HandleFunc("/hc", func(w http.ResponseWriter, r *http.Request) {
		response := w
		if w == nil {
			fmt.Fprintf(response, "hc is not working!")
		} else {
			fmt.Fprintf(response, "hc is working...")
		}
	})
}
