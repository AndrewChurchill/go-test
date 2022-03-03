package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bytesWritten, err := fmt.Fprintf(w, "Hello, New York!!!")
		if err != nil {
			fmt.Println(err)
		}

		log.Printf("Wrote %d bytes", bytesWritten)
	})

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
		return
	}

	log.Println("Starting server on port: " + port + "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
