package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hello world"))
		if err != nil {
			log.Println("write response failed:", err)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
