package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bytesWritten, err := fmt.Fprintf(w, "Hello, New York!!!")
		if err != nil {
			fmt.Println(err)
		}

		log.Printf("Wrote %d bytes", bytesWritten)
	})

	log.Println("Starting server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
