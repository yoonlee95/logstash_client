package main

import (
    "net/http"
    "log"
    "fmt"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/hokaccha/go-prettyjson"
)

var (
    req_num = 0
)

func LogStashCallback(w http.ResponseWriter, r *http.Request) {
    fmt.Println("-----LogStash Endpoint-----")
    fmt.Printf("Request Number %d\n", req_num);
    req_num += 1

    // Decode Json
    // m := map[string]string{}
    var m map[string]*json.RawMessage
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&m)
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()
    
    // Display data 
    fmt.Println("----Data ----")
    s, _ := prettyjson.Marshal(m)
    log.Print(string(s))

}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/logstash_endpoint", LogStashCallback)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":20000", r))
}
