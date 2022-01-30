package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/postmessage", postDataToConsoleHandler).Methods("POST")
	staticFileDir := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDir))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8081", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listening on 8081")
}

func postDataToConsoleHandler(w http.ResponseWriter, r *http.Request) {
	var posts PostMessage
	// requestBody, _ := ioutil.ReadAll(r.Body)
	json.NewDecoder(r.Body).Decode(&posts)
	// json.Unmarshal(requestBody, &posts)
	fmt.Println(posts.message)
	// fmt.Printf("message: %s", requestBody)
}
