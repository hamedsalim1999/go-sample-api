package main

import (
        "io"
        "log"
		"os"
        "net/http"
)
func main() {
		logger := log.New(os.Stdout, "http: ", log.LstdFlags)
		logger.Println("Server is starting...")
        http.HandleFunc("/", Tmp)
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                log.Fatal(err)
        }
}

func Tmp(w http.ResponseWriter, r *http.Request) {
		logger := log.New(os.Stdout, "http: ", log.LstdFlags)
        io.WriteString(w, "42")
		logger.Println("Status:",200)
}