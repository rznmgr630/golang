package main;

import "fmt";

func main(){
	slice1 :=[]int{1,2,3,4,5}
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))

	// ========================== Create a Slice From an Array =========================
	var myArray=[6]int{1,2,3,4,5,6}
	mySlice :=myArray[1:4];

	fmt.Printf("myslice = %v\n", mySlice)
  fmt.Printf("length = %d\n", len(mySlice))
  fmt.Printf("capacity = %d\n", cap(mySlice))


	// =========================== Create a Slice With The make() Function ==============
	// make(type,length, capacity)
	// if capacity is not given then its value will be equal to the length
	myslice1 := make([]bool, 5, 5)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))


	// =========================== To access the value of slice =========================
	fmt.Println(myslice1[0])


	// =========================== To change the value of slice =========================
	myslice1[0]=true;
	fmt.Println(myslice1)

	// =========================== Append element to the slice ==========================
	fmt.Println("New---------------------------")
	myslice1=append(myslice1,false,true)
	fmt.Println(myslice1,len(myslice1),cap(myslice1))

	// =========================== REMOVE ELEMENT FROM THE SLICE =========================
	myslice1=append(myslice1[:2],myslice1[2+1:]...);
	fmt.Println(myslice1,len(myslice1),cap(myslice1))
}