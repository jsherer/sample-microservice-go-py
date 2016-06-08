package main

import (
    "fmt"
    "log"
    "time"
    "net/http"
)

func main(){
    http.HandleFunc("/hello", handleHello)
    http.HandleFunc("/sleep", handleSleep)
    fmt.Println("serving on http://localhost:8000/")
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
    log.Println("serving", req.URL)
    fmt.Fprintln(w, "Hello, World!")
}

func handleSleep(w http.ResponseWriter, req *http.Request) {
    time.Sleep(10 * time.Second)
    fmt.Fprintln(w, "Done")
}
