package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	fmt.Println("Starting server at port 80")
	log.Println("Starting http://localhost:80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}
