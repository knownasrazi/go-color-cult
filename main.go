package main

import "net/http"

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("go-color-cult - Colors, cult-grade Go."))
  })
  http.ListenAndServe(":8080", nil)
}
