package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/Daksh-Jain08"

func main() {
	fmt.Println("Web requests in golang")

	response, err := http.Get(url)
	checkErr(err)

	fmt.Printf("The type of the response is %T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	checkErr(err)

	content := string(databytes)
	fmt.Println(content)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
