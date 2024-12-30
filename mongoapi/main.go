package main

import (
	"fmt"
	"mongoapi/routers"
	"net/http"
)

func main() {
	fmt.Println("hello world")

	r := routers.Routers()
	fmt.Println("server started on port 3000")
	http.ListenAndServe(":3000", r)
}
