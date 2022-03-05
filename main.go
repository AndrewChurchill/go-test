package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/hidden-message", handleHiddenMessage)

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

func handleRoot(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	bytesWritten, err := fmt.Fprintf(w, "Hello, New York!!!")
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Wrote %d bytes", bytesWritten)
}

func handleHiddenMessage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	bytesWritten, err := fmt.Fprintf(w, "YOU FOUND IT!!! CONGRATS!!!")
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Wrote %d bytes", bytesWritten)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
