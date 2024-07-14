package main;

import "fmt";

func main(){
	// =============================== Array with fixed length ==================================
	var items =[3]int{1,2,3};
	fmt.Println(items)

	// =============================== Array without fixed length ===============================
	var items1=[...]int{1,2,3,4,5,6,7}
	var items2=[...]string{"Raju","Raja","Ram"}
	fmt.Println(items1,items2);

	// =============================== Accessing the element ====================================
	fmt.Println(items[0])

	// =============================== Changing the element =====================================
	items[0]=100;
	fmt.Println(items)

	// =============================== Array intitalization =====================================
	arr1 :=[2]int{}  // not initalized
	arr2 :=[2]int{1}  // partially inititalized
	arr3 :=[2]int{1,2}  // fully initialized

	fmt.Println(arr1,arr2,arr3)

	// =============================== Initalized only specfic postion element ==================
	arr4 := [4]int{0:1,3:23}
	fmt.Println(arr4)

	// =============================== To get the length of an array ============================
	fmt.Println(len(arr1))
}