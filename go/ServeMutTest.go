package main

import (
    "net/http"
    "io"
)

type b struct{}

func (*b) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello")
}
func main() {
    mux := http.NewServeMux()
    mux.Handle("/h", &b{})
    http.ListenAndServe(":8080", mux)
}
