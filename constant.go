package main

import "fmt"

const PI =3.14;

func main(){
	// typed constant
	const A int=1;

	// untyped constant
	const B=45;

	fmt.Println("Value of PI is",PI);
	fmt.Println("Value of A is",A);
	fmt.Println("Value of B is",B);

	// Multiple constant declaration
	const (
		C =123
		D int =34
		E ="Hello"
	)
	fmt.Println(C,D,E);
}
