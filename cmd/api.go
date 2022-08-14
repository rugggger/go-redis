package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}

func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}
	err = Put(key, string(value))
	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func serve() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloGoHandler)
	r.HandleFunc("/v1/{key}", keyValuePutHandler).Methods("PUT")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}

}
