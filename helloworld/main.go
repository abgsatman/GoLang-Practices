package main

import (
	"fmt"
	"net/http"
)

func main() {
	a := 2
	fmt.Println("a: ", a)

	//index()
	//hc()

	printLib()

	//http.ListenAndServe(":80", nil)
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

func printLib() {
	fmt.Println("printing with a line")

	i := 2
	s := "golang"
	f := 3.14

	fmt.Printf("integer: %d\n", i)
	fmt.Printf("string: %s\n", s)
	fmt.Printf("float: %f\n", f)

	fmt.Println("------------------------------------------")

	fmt.Printf("the type of the first one is: %[1]T\n", i)
	fmt.Printf("the type of the second one is: %[1]T\n", s)
	fmt.Printf("the type of the last one is: %[1]T\n", f)

	fmt.Println("------------------------------------------")

	fmt.Printf("the value of the first one is: %-4v\n", i)
	fmt.Printf("the value of the second one is: %-4v\n", s)
	fmt.Printf("the value of the last one is: %-4v\n", f)
}
