package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func checker(err error) {
	if err != nil {
		panic(err)
	}
}

const url = "http://localhost:8000/get"

func main() {

	// callGetMethoc()
	postMethod()

}

func postMethod() {
	const insideUrl = "http://localhost:8000/post"
	reqBody := strings.NewReader(`
	{
	"name": "manoj bhatta",
	"age": 52,
	"work": "software deveoper"
	}
	`)
	res, err := http.Post(insideUrl, "application/json", reqBody)
	checker(err)
	defer res.Body.Close()
	if res.StatusCode == 200 {
		fmt.Println("post response is 200 so successfully gotten")
	} else {
		fmt.Println("post response is not  200 so something went wrong")
	}

	content, err := io.ReadAll(res.Body)

	fmt.Println("the  actual byte is ", content)

	// fmt.Println("the actual data is ", string(content))
	var respo_content Respo

	err = json.Unmarshal(content, &respo_content)
	checker(err)
	fmt.Println(respo_content.Name)
	fmt.Println(respo_content.Age)
	fmt.Println(respo_content.Work)

}

type Heading struct {
	Message string `json:"message"`
}

type Respo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Work string `json:"work"`
}

func callGetMethoc() {
	resp, err := http.Get(url)
	checker(err)
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	checker(err)
	var processedData Heading
	err = json.Unmarshal(data, &processedData)
	checker(err)
	fmt.Println(processedData.Message)

}
