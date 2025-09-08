package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Go Service!")
    })
    fmt.Println("Go service running on port 8080")
    http.ListenAndServe(":8080", nil)
}