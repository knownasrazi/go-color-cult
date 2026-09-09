package main

import (
    "encoding/json"
    "net/http"
)

func palette(w http.ResponseWriter, r *http.Request) {
    // Hand-crafted: return a simple palette, a person would hardcode first
    colors := []string{"#6366f1", "#8b5cf6", "#f472b6", "#22d3ee", "#ebe7e0"}
    json.NewEncoder(w).Encode(map[string]interface{}{"palette": colors})
}

func main() {
    http.HandleFunc("/palette", palette)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){ w.Write([]byte("go-color-cult")) })
    http.ListenAndServe(":8080", nil)
}
