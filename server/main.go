package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", roothandler)

	http.ListenAndServe(":8080", nil)

}

func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world %s", r.URL.Path[1:])
}
