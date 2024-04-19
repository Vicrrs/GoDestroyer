package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

var statusCode = 200

func signalHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(statusCode)
}

func updateStatusHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if newCode, err := strconv.Atoi(code); err == nil {
		statusCode = newCode
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func destroyProjectHandler(w http.ResponseWriter, r *http.Request) {
	err := os.RemoveAll("/home/tkroza/Documentos/fast_api_test")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/signal", signalHandler)
	http.HandleFunc("/update-status", updateStatusHandler)
	http.HandleFunc("/destroy-project", destroyProjectHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
