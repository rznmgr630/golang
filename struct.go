package main;

import "fmt";

type Person struct {
	id int
	name string
	address string
	age int
}

func main(){
	var person1 Person;
	var person2 Person;

	 // Pers1 specification
	 person1.id = 1
	 person1.name = "Hege"
	 person1.address = "Kalanki"
	 person1.age = 45

	 // Person2 specification
	 person2.id = 2
	 person2.name = "Raju"
	 person2.address = "Kalimati"
	 person2.age = 23

	 // Person 3 specification
	 person3 :=Person{3,"Kaka","Kathmandu",34};

	 printDetail(person1)
	 printDetail(person2)
	 printDetail(person3)
	fmt.Printf("person 3 details is %+v\n",person3)
}


func printDetail(person Person){
	fmt.Println("ID",person.id)
	fmt.Println("Name",person.name)
	fmt.Println("Age",person.age)
	fmt.Println("Address",person.address)
}