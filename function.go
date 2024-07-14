package main;

import "fmt";

// ============================== BASIC FUNCTION ====================================== 
func add(){
	fmt.Println("The sum of 5 and 6 is", 5+6)
}

// ============================== PARAMETER ==========================================
func subtract(first int, second int){
	fmt.Printf("The subtract of %d and %v is %v\n",first,second, first-second);
}

// ============================== RETURN =============================================
func multiply(first int, second int) int {
	return first*second;
}

// ============================== NAMED RETURN ========================================
func divide(first int, second int) (result int) {
	result = first/second;
	// return OR
	return result;
}

// ============================== Multiple Return Values ==============================
func multipleReturnValue(first int, second int) (result int, name string) {
	result = first+second;
	name="Rajan Midun Magar";
	return;
}


// ============================== Function that takes any number of input =============
func canTakeAnyValue(values ...int) int{
	total :=0;
	for	_,value := range values{
		total+=value;
	}
	return total;
}

func main(){
	add();
	subtract(4,3)
	fmt.Println("The multiply of 4 and 5 is",multiply(4,5));
	fmt.Println("The division of 4 and 2 is",divide(4,2));
	sum,name :=multipleReturnValue(4,2);
	_,name2 :=multipleReturnValue(4,2); //here _ is used to omit the value
	fmt.Println("Multiple return value",sum,name,name2);

	fmt.Println("The sum is ",canTakeAnyValue(2,3,2,3,2,3));
}