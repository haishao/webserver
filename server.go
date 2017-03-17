package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type Input struct {
    Id      string
    Name    string
}

type Output struct {
    Id      string
    Message string
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    var input Input
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        panic(err)
    }

    output := Output{Id: input.Id, Message: "results here ..."}

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(output); err != nil {
        panic(err)
    }
}
