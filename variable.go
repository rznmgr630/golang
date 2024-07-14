package main

import "fmt"

func main(){
	// ============================= string ============================
	var firstName string="Rajan";
	var lastName string="Midun Magar";
	var address string;
	address="Dhungeadda";

	fmt.Println("Your name is",firstName,lastName);
	fmt.Println("Your address is",address);

	country := "Nepal";
	fmt.Println("Your country is",country)

	// ============================ integer ==========================
	var age int =12;
	age2 :=14;
	var age3 int;
	age3=45;

	fmt.Println("Age is", age,age2,age3)


	// ========================== float===============================
	var income float32 =12.09;
	income2 :=14;
	var income3 float32;
	income3=45;

	fmt.Println("Income is", income,income2,income3);

	// Multiple variable declaration
	var a,b,c,d int=1,2,3,4;
	var e, f =12,14;
	g, h := 34,45;
	
	fmt.Println(a,b,c,d,e,f,g,h)


	// =========================== Variable declaration in block ==========
	var (
		aa int 
		bb int =23
		cc string="Kaka"
	)
	fmt.Println(aa,bb,cc)


	// =========================== Boolean =========================
	var b1 bool =true;
	var b2=false;
	var b3 bool;
	b4 := true;

	fmt.Println(b1,b2,b3,b4)
}