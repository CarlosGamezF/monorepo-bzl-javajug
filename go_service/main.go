// go_service/main.go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Go!")
}

func main() {
    fmt.Println("Ready :8081/hello")
    http.HandleFunc("/hello", helloHandler)
    http.ListenAndServe(":8081", nil)

}