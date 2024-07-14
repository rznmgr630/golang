package main;

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
)

func main(){
	fmt.Println("Read and write to the file");

	file,error := os.Create("./src.txt");
	checkError(error);

	length, error := io.WriteString(file,"Hello from rajan");
	checkError(error);


	fmt.Println("length is ",length)

	readFile("./src.txt")

	defer file.Close();
}

func readFile(filename string){
	databyte,error	:= ioutil.ReadFile(filename);
	checkError(error);

	fmt.Println(string(databyte))
}

func checkError(err error){
	if err != nil {
		// panic will shot down execution of the program and show you the error
		panic(err);
	}
}