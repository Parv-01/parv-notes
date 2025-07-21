package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go arrays")
	/*Arrays*/
	//arraysgo()
	/*Arrays Operations*/
	//arrayops()
	/*Arrays with Keyed ELements*/
	//arryskeydelements()
}

// func arraysgo(){
// 	/*Array :- fixed length,indexable type,stores collection of item
// 	Stores elements in contiguous memory locations
// 	Every element in slice and array should be of same type
// 	length and elements type determines type of array
// 	length belongs to array type & defined at compile time
// 	*/
// 	acc := [3]int{2,3,4}
// 	fmt.Println(acc)
// 	fmt.Printf("%T %#v\n",acc,acc) //#will tell the type alongwith length of array
// 	//Possible to initialise some array
// 	af:=[3]float64{0.9,0.8}
// 	fmt.Println(af)
// 	//How to get the length of array [...] -> ellipse operator
// 	ac:=[...]int{5,8}
// 	fmt.Printf("%d\n",len(ac))//len tells length of array
// 	//For multiline array the last , is mandatory
// 	ag:=[...]int{
// 		1,
// 		2,
// 		3,
// 		6, //this , is  mandatory
// 	}
// 	fmt.Println(ag)
// }

// func arrayops(){
// 	num := [3]int{1,3,4}
// 	fmt.Println(num)
// 	num[0]=8 //to change value at particular element index
// 	fmt.Println(num)
// 	//TO print the ararys
// 	for i,v :=range num{
// 		fmt.Println(i,v)
// 	}
// 	for i:=0;i<len(num);i++{
// 		fmt.Println(i,num[i])
// 	}
// 	//Array of arrays :- Multidimensionl Arrays
// 	bal:=[...][4]int{//2 lines 4  cols
// 		{5,6,7},
// 		[4]int{1,2,3,4},
// 	}
// 	fmt.Println(bal)
// 	//Equality of arrays :-  said to be equal if size and elements are same
// 	m:=[3]int{1,2,3}
// 	n:=m
// 	fmt.Println("n==m",n==m)
// 	m[0]=-2
// 	fmt.Println("n==m",n==m)
// }

// func arryskeydelements() {
// 	/*Index start from 0
// 	unkeyed element gets the index from last keyed element
// 	keys are int
// 	*/
// 	num := [3]int{ //Keyed elements can be in any order
// 		1: 44,
// 		2: 33,
// 		0: 66,
// 	}
// 	fmt.Println(num)
// 	num1 := [3]int{
// 		2: 33,
// 	}
// 	fmt.Println(num1)
// 	num2 := [...]string{
// 		6: "fdfcd x",
// 	}
// 	fmt.Println(num2, len(num2))
// 	num3 := [...]string{
// 		6:         "fdfcdx",
// 		"aaaaaaa", //index 7
// 		0:         "rrrr", //elements 8
// 	}
// 	fmt.Printf("%#v %d", num3, len(num3))
// }


