package main

import (
	"fmt"
)

func GoSlices() {
	fmt.Println("This file for Slices")
	/*Difference Between arrays and slices */
	// arrayVsSlice()
	/*Slices Declaration*/
	//declareSlice()
	/*Slice Comparison*/
	//compareSlices()
	/*Copying Slices*/
	copySlice()

}

/*Differrence between arrays and slices*/
// func arrayVsSlice(){
// 	/*Arrays ->fixed length
// 	Slice -> Dynamic Length
// 	//////////////////////////////////////////////////////////
// 	Arrays : length is part of its type,defined at compile time and cant be changed
// 	Slice : length not part of it type ,belongs to runtime
// 	/////////////////////////////////////////////////////////
// 	Array : by default uninitialised array has all elements 0
// 	Slice : by default unitialised slice is equal to nil
// 	////////////////////////////////////////////////////
// 	---------------SIMILARITIES------------------------
// 	Both Arrays and slices only contains same type of elements
// 	Keyed slice can be created like keyed arrays
// 	*/
// }

// func declareSlice(){
// 	var cities []string
// 	fmt.Println("It should be nill values ",cities==nil)
// 	fmt.Printf("%T %s %d One more methos is like this %#v\n",cities,cities,len(cities),cities)
// 	// cities[0]="abc" //gives error of index out of range
// 	//Creating slice with elements
// 	nums:=[]int{11,12,13,14} //Slice literal
// 	fmt.Println(nums)
// 	//Another way to create slice using make which takes length and optional cap as arguments
// 	numb:= make([]int,2)
// 	numb[0]=34
// 	fmt.Println(numb)
// 	//Creating a new type
// 	type fr []string
// 	frrr :=fr{"aaa","grfvf"}
// 	fmt.Println("heyyyeyeyey ",frrr[0])
// 	//Modifying elements using indexing
// 	frrr[0]="cdbvhvhgv"
// 	fmt.Println("heyyyeyeyey ",frrr[0])
// 	//To see slice using for loop
// 	for index,value := range nums{
// 		fmt.Println(index,value)
// 	}
// 	//NOTE :- slices of same type can be assigned to eahc other,if they have been erased then they should have same length
// 	var n []int
// 	n=nums
// 	fmt.Println(n)
// }

// func compareSlices() {
// 	// var n []int
// 	// fmt.Println("n==nil ",n==nil)
// 	// //Slice declare but not equal to nil
// 	// m:=[]int{}
// 	// fmt.Println("m==nil ",m==nil)
// 	//Cant compare directly using ==
// 	a, b := []int{1, 2}, []int{1, 2}
// 	// fmt.Println(a==b) //this is error conditionn,slice can only be compared to nil
// 	//How to compare slices ->using for loop
// 	var eq = true
// 	// b[1]=7
// 	// for i,val:=range a{
// 	// 	if val!=b[i]{
// 	// 		eq = false
// 	// 		fmt.Println("Not Equals eq == ",eq) //printing eq statement
// 	// 		break //breaking out of loop
// 	// 	}
// 	// }
// 	//but this will create problem if we a=nil and give wrong answer as below
// 	a = nil
// 	// for i,val:=range a{
// 	// 	if val!=b[i]{
// 	// 		eq = false
// 	// 		break //breaking out of loop
// 	// 	}
// 	// }
// 	// if eq{
// 	// 	fmt.Println("equals")
// 	// 	}else{
// 	// 	fmt.Println("Not Equals eq == ",eq) //printing eq statement
// 	// }
// 	/*To solve the error We will put one MOST IMPORTANT CONDITION WHEN COMPARING Slices*/
// 	if len(a) != len(b) {
// 		fmt.Println("not equal due to length")
// 	} else {
// 		for i, val := range a {
// 			if val != b[i] {
// 				eq = false
// 				break //breaking out of loop
// 			}
// 		}
// 		if eq {
// 			fmt.Println("equals")
// 		} else {
// 			fmt.Println("Not Equals eq == ", eq) //printing eq statement
// 		}
// 	}
// }

func copySlice() {

}
