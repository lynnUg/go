package main

import (
    "fmt"
    "net/http"
)

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (s String) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, "%v", string(s))
}

func (s *Struct) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, "%v",s)
}

func main() {
    // your http.Handle calls here
    http.ListenAndServe("localhost:4000", nil)
    http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
}
