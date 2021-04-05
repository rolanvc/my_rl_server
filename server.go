package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", staticFileServer)
	r.HandleFunc("/test/", handler)

	fmt.Printf("Starting RL Server at port:8080\n")
	http.ListenAndServe(":8080", r)

}
func staticFileServer(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello, World!"))
}
