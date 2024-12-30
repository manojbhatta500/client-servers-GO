package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const jsonUrl = "https://jsonplaceholder.typicode.com/"

type todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	resp, err := http.Get(jsonUrl + "todos/1")
	if err != nil {
		fmt.Println("something wentwrong while  fetching data from this url : ", jsonUrl)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		var goData todo

		err = json.Unmarshal(data, &goData)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		fmt.Println(goData)

	}
}
