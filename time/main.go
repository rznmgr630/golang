package main;

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Hello world");

	currentTime := time.Now();
	fmt.Println(currentTime);

	// formattting the date
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))

	// TO CREATE A NEW DATE
	createdAt	:= time.Date(1996, time.October, 11, 06, 06, 0, 0, time.UTC);
	fmt.Println(createdAt)
	fmt.Println(createdAt.Format("01-02-2006 15:04:05 Monday"))
}