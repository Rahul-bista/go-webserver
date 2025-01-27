package main

import (
	"fmt"
	"net/http"
)

func handleForm(w http.ResponseWriter, r *http.Request){
	if err:= r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful!\n")
	name := r.FormValue("name")
	age := r.FormValue("age")
	fmt.Fprintf(w, "Name = %v \nAge = %v", name, age)
}

func handleHello(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w, "Method not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleHello)


	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("error")
	}
}