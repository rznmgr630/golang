package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Reading data from the json server")
	getTodos()
	addTodo()
	formData()
}

func getTodos() {
	response, error := http.Get("https://jsonplaceholder.typicode.com/todos")
	if error != nil {
		panic(error)
	}

	fmt.Printf("Response type is %T\n", response)

	defer response.Body.Close()

	databyte, error := ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}

	fmt.Println(string(databyte))
}

func addTodo() {
	requestBody := strings.NewReader(`
	{
		"title": "foo",
    "body": "bar",
    "userId": 1
	}
	`)
	response, error := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", requestBody)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	databyte, error := ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}

	fmt.Println(string(databyte))
}

func formData() {
	data := url.Values{}
	data.Add("title", "hello")
	data.Add("body", "hello")
	data.Add("userId", "1")

	response, error := http.PostForm("https://jsonplaceholder.typicode.com/posts", data)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	databyte, error := ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}

	fmt.Println(string(databyte))
}
