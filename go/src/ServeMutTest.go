package main

import (
	"fmt"
	"io"
	"net/http"
)

type b struct{}

func (*b) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "hello")
	w.WriteHeader(404)
	fmt.Println(w.Header().Get("Status Code"))
	io.WriteString(w, fmt.Sprintf("%s", r.Header))
	io.WriteString(w, fmt.Sprintf("%s", w.Header().Get("Status Code")))
	io.WriteString(w, fmt.Sprintf("%s", w))
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/h", &b{})
	http.ListenAndServe(":18787", mux)
}
