package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8088", nil)
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("go-server is starting...")
	http.HandleFunc("/", Tmp)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Tmp(w http.ResponseWriter, r *http.Request) {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	io.WriteString(w, "42")
	logger.Println("Status:", 200)
}
