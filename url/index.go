package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Hello form url file")

	result, error := url.Parse("https://www.google.com?search=rajan")

	if error != nil {
		panic(error)
	}

	fmt.Println("Scheme", result.Scheme)
	fmt.Println("Host", result.Host)
	fmt.Println("Path", result.Path)
	fmt.Println("Port", result.Port())
	fmt.Println("RawQuery", result.RawQuery)

	queries := result.Query()
	fmt.Println(queries["search"])
}
