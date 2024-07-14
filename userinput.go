package main;

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Welcome to rating app!")

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter your rating: ");

	// comma ok syntax
	input ,_ := reader.ReadString('\n');
	fmt.Println("Thanks for the rating",input) 
	fmt.Printf("Type of your rating is %T",input) 

	numberRating,err :=strconv.ParseFloat(strings.TrimSpace(input),64);

	if err != nil {
		fmt.Println(err)
	}	else {
		fmt.Println("Added 1 to rating",numberRating+1)
	}
}