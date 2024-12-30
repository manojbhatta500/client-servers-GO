package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "https://www.dadabhagwan.org/path-to-happiness/humanity/help-others-the-purpose-of-life/?gad_source=1&gclid=CjwKCAiAjp-7BhBZEiwAmh9rBROnl226_PWKcgc8vyqHPwmPM5xrkTM8VQAlpRLa0dY8AUMrg7mVWRoC18wQAvD_BwE"

func main() {

	fmt.Println("hello welcome to file getter program ")

	response, err := http.Get(url)

	errorChecker(err)

	data, err := io.ReadAll(response.Body)

	defer response.Body.Close()

	errorChecker(err)

	fmt.Println("the actual data is ", string(data))
	var filename string = "output.html"
	file, err := os.Create(filename)

	_, err = file.Write(data)

	fmt.Println("successfully written to file ", file.Name())
	// fmt.Println(n)

	defer file.Close()

}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
