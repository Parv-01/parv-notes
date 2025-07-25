package main

import (
	"fmt"
	//"unsafe" //Used for finding the size of arrays and slices
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
	// copyAppendSlice()
	/*Slice Expression*/
	// sliceExpression()
	/*Slice Backing(underlying) Array*/
	// sliceBacking()
	/*Append Length and Capacity in Depth*/
	// appendfunc()
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

// func copyAppendSlice() {
// 	nums:=[]int{2,3}
// 	//append doesnt modify initial slice but returns a new slice 
// 	nums= append(nums,10)
// 	fmt.Println(nums)
// 	//append function can append more values at once to slice in last 
// 	nums =append(nums,23,45,55)
// 	fmt.Println(nums)
// 	//We can append other slice in other slice 
// 	n:=[]int{0,0,0}
// 	nums=append(nums,n...) //will append the whole n slice to nums in last 
// 	fmt.Println(nums)
// 	//copy function copies elements of one slice to another 
// 	//If the slices have different length then minimum of length of slices 
// 	src:=[]int{10,20,30}
// 	dst:=make([]int,len(src))
// 	nn:=copy(dst,src)
// 	fmt.Println(src,dst,nn)	
// 	dst=make([]int,2)
// 	nn=copy(dst,src)
// 	fmt.Println(src,dst,nn)	
// 	//Note:=copy function not always copy elements if nothing is present then nothing is copied 
// 	dst=make([]int,0)
// 	nn=copy(dst,src)
// 	fmt.Println(src,dst,nn)
// }

// func sliceExpression(){
// 	a:=[5]int{1,2,3,4,5}
// 	//a[start:stop] -> will return element from start index to excluding stop index 
// 	b:=a[0:3]///will return 1,2,3 ->slicing an array return slice 
// 	fmt.Printf("%v, %T\n",b,b)
// 	//Mising start says 0 and stop by len(slice)
// 	b=a[:2]
// 	fmt.Println(b)
// 	b=a[2:]
// 	fmt.Println(b)
// 	b=a[:]//new slice equal to a
// 	fmt.Println(b)
// 	b=append(a[:3],100)
// 	fmt.Println(b)
// 	b=append(a[:3],200)//It pverwrites the last element from 100 to 200
// 	fmt.Println(b)
// }

// func sliceBacking(){
// 	/*Behind creating slice go creates array called Backing Array,
// 	Backing array stores elements not slice 
// 	Go implements slice as data structure called slice header
// 	Slice header contains 3 fields :- 
// 	1)address of backing array(ptr)
// 	2)length of slice -> len() returns it
// 	3)capacity of slice -> cap() returns it
// 	Slice header is runtime representation of slice 
// 	NOTE:- nil slice doesnt have backing array(all fiels of slice header 0)
// 	The slice expression doesnt create new backing array 
// 	*/
// 	// s1:=[]int{10,20,30,40,50}
// 	// s3,s4:=s1[0:2],s1[1:3]
// 	// s3[1]=600 //The baking array of s1,s3,s4 is changed as well
// 	// fmt.Println(s1,s3,s4)
// 	// arr:=[4]int{10,20,30,40}//array
// 	// slice1,slice2:=arr[0:2],arr[1:3]
// 	// arr[1]=300 //Modifying array also modify slices 
// 	// fmt.Println(arr,slice1,slice2)
// 	// //To create a complete new slice from existing use append 
// 	// s2:=[]int{}
// 	// s2=append(s2,s1[0:2]...)//Remember ... else it will create error
// 	// s1[0]=200//This doesnt modify new slice as both have different backing array
// 	// fmt.Println(s1,s2)
// 	// s5:=s1[0:3]
// 	// fmt.Println(len(s5),cap(s5))//Capacity will be of s1 and length will be of s5 as both have same backing array
// 	// s5=s1[2:]//The cap will be changes as in backing array the address of backing array(ptr) is the address of 1 element of slice
// 	// fmt.Println(len(s5),cap(s5))
// 	// //Slicing operations are cheaper then arrays as they have same backing arrays 
// 	// fmt.Printf("%p %p %p\n",&s1,&s3,&s2)
// 	// s5[0]=20000
// 	// fmt.Println(s5,s1)//due to same backing array the element change in s5 will also be reflected in s1
// 	// a:=[5]int{1,2,3,4,5}
// 	// s:=[]int{1,2,3,4,5}
// 	// fmt.Printf("Arrays size in bytes %d \n",unsafe.Sizeof(a))
// 	// fmt.Printf("slices size in bytes %d \n",unsafe.Sizeof(s))//slices takees less memory than arrays 
// }

// func appendfunc(){
// 	// var nums []int
// 	// fmt.Printf("val: %#v, Length: %d,Capacity: %d\n",nums,len(nums),cap(nums))
// 	// nums=append(nums,1,2)
// 	// /*Capacity of slice(0) is length of backing arrays and length of arrays fixed in golang 
// 	// Remember elements of slice stores in backing arrays and not slices 
// 	// So when slice capacity is full it makes new backing array*/
// 	// fmt.Printf("val: %#v, Length: %d,Capacity: %d\n",nums,len(nums),cap(nums))
// 	// /*If we add more inside this then it will create one more new backing array*/
// 	// nums=append(nums,3)
// 	// fmt.Printf("val: %#v, Length: %d,Capacity: %d\n",nums,len(nums),cap(nums))//it adds 1 extra for future append
// 	// nums=append(nums,4)
// 	// fmt.Printf("val: %#v, Length: %d,Capacity: %d\n",nums,len(nums),cap(nums))//this time it didnt create new backing array 
// 	// nums=append(nums,5)
// 	// fmt.Printf("val: %#v, Length: %d,Capacity: %d\n",nums,len(nums),cap(nums))//but this increase the backing arrays capacity exponentially
// 	//
// 	letters := []string{"a","b","c","d","e","f"}
// 	letters=append(letters[0:1],"x","y")
// 	fmt.Printf("val: %#v, Length: %d,Capacity: %d\n",letters,len(letters),cap(letters))//still the same backing array of previous letters
// 	/*ERROR!!!!!!*/
// 	// //fmt.Println(letters[4])//cant access after its length been append so give error
// 	//but this will work using slice because slices see the whole backing arrays
// 	fmt.Println(letters[3:6])
// }