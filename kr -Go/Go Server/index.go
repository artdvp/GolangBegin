package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "go check Test")
		fmt.Println("log")
	})
	http.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Me Go")
	})
	http.HandleFunc("/save", saveServ)
	http.HandleFunc("/user", user)
	http.ListenAndServe(":8066", nil) // nil => null
}

func saveServ(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "save")
	fmt.Println("log save")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user")
	fmt.Println("log user")
}
