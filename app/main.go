package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("launching server")
	port, _ := ":", os.Getenv("SERVER_PORT")
	e := http.ListenAndServe(port, nil)
	if e != nil {
		log.Fatal("server launche error \n", e)
	}
}
