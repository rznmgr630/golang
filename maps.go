package main;

import "fmt";

func main(){
	languages := make(map[string]string);
	languages["JS"]="Javascript";
	languages["PY"]="Python";
	languages["JAVA"]="Java";

	fmt.Println(languages)
	delete(languages, "PY")
	fmt.Println(languages)

	for key,value:=range languages{
		fmt.Println(key, "-",value)
	}
}