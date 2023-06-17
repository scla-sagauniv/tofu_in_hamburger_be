package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("launching server")
	e := http.ListenAndServe(":8080", nil)
	if e != nil {
		log.Fatal("server launche error \n", e)
	}
}
