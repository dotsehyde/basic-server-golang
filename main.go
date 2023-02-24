package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//create fileServer := to get html files
	fileServe := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServe)
	http.HandleFunc("/form", formFunc)
	http.HandleFunc("/hello", helloFunc)
	fmt.Println("Server running on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}

}

// Hello endpoint Handler
func helloFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Invalid Path", http.StatusBadRequest)
	}
	fmt.Fprint(w, "hello!")
}

// Form endpoint Handler
func formFunc(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful\n")
	//Get Form values
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}
