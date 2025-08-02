package api

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type StatusResponse struct {
    Message string `json:"message"`
}

func StartServer() {
    http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
        res := StatusResponse{Message: "OmniForge is running"}
        json.NewEncoder(w).Encode(res)
    })

    http.HandleFunc("/spawn", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Spawning new multiverse environment...")
    })

    fmt.Println("API server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}