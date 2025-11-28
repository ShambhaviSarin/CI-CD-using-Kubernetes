package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/ShambhaviSarin/CI-CD-using-Kubernetes/src/greet"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    msg := greet.Hello(name)
    fmt.Fprintln(w, msg)
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    addr := fmt.Sprintf(":%s", port)
    log.Printf("starting server on %s", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}
