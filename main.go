package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/google/uuid"

    "github.com/ShambhaviSarin/CI-CD-using-Kubernetes/src/greet"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // generate a small request id using google/uuid
    reqID := uuid.New().String()
    w.Header().Set("X-Request-ID", reqID)

    name := r.URL.Query().Get("name")
    msg := greet.Hello(name)
    fmt.Fprintf(w, "%s\n", msg)
    log.Printf("request-id=%s path=%s name=%s", reqID, r.URL.Path, name)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello from Kubernetes CI/CD!")
    })

    http.ListenAndServe(":8080", nil)
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
