//Initial Hello World Program in Go
// package main
// import (
// 	"fmt"
// )
// func main(){
// 	fmt.Println("Hello, World!")
// }

// Variables and Declarations
package main

//to deal with very very large numbers use math/big :- it implements arbitrary-precision arithmetic(big no)
//foolwing are supported :-  Int -> signed int , Rat -> rational no,Float -> floating point no ;zzero val is 0 for all 3

import (
	"fmt"
	//"math"
)
func main(){
	fmt.Println("Hallo Parv,Ich lebe in India!!")

	// //2 ways to declare var
	// //1 type :- var name type = value | keyword identifier type = value
	// var name string = "Parv"
	// fmt.Println(name)
	// var tr = 2
	// //fmt.Println(tr) //if we comment this line then it will shows compile type error by saying "tr declared but not used"
	// //To use var without using it we can use _
	// _ =tr //this will not through any kind of error
	// //2 type :- var name = value | keyword identifier = value
	// name2 := "Doe"
	// fmt.Println(name2)
	// //cant use for already defined variable

	// // Multiple Declarations
	// a,b:="nnkn",45
	// fmt.Println(a, b)
	// //a,b:="v,,d",435 //compile time erro :- this short declaration can only be used if at left hand side we have atleast one new var to declare 
	// a,t:="v,,d",435 //this is fine as we are declaring a new variable t
	// fmt.Println(a, t)

	// //Multiple Declarations with var
	// var (
	// 	x int
	// 	y string
	// 	z float64
	// )
	// fmt.Println(x,y,z) //this will print 0 "" 0 as default values of int, string and float64
	// //We use short declaration when we know the initial val of var and normal otherwise or better readability

	// //Multiple assignment 
	// var i,j int
	// i,j=5,6
	// _,_ = i, j //this will not throw any error as we are using _ to ignore the values(blank identifier)
	// j,i=i,j //this will swap the values of i and j
	// fmt.Println(i, j) //this will print 6 5

	// //Types and Zero values :- in go its necessay to declare or type or to get automatically the type of var if not provided
	// var a int=2
	// var b float64=4.23
	// /*a=b//it will throw compile time error as we are trying to assign float64 to int,The type of var once declared cannot be changed
	// To fix this we can use type conversion*/
	// a = int(b) //this will convert float64 to int and assign it to a
	// fmt.Println(a,b)
	/*Go is strongly typed language, so we cannot assign a value of one type to a variable of another type without explicit conversion
	Zero values are the default values of a variable when it is declared but not initialized,in go theres no uninitialized variable
	zero values are: numeric : 0, string : "", bool : false, pointer : nil, slice : nil, map : nil, channel : nil, function : nil, interface : nil*/

	//Naming Conventions
	/*
	Name begins with letter,underscore(_)
	Case Matters,
	keywords(25 in no) cant be used 
	*/
	//Best Practices
	/*
	Use CamelCase for multi-word identifiers instead of snake_case
	Use meaningful names
	Use short names for small scopes and longer names for larger scopes
	Use all caps for acronyms (e.g., HTTP, URL)
	An upper case first letter mostly goes in package functions(it will be exported in otther packages)so use first upper case letter for exported functions and variables in other packages and files
	NOTE:- Go doesnt provide automatic support for getters and setters like other languages, so we have to write them manually if needed
	One method interfaces are named by the method name + an er suffix (e.g., Reader, Writer, Stringer)
	*/

	// //FMT Package :- std library for formatting and printing
	// //1 Println
	// fmt.Println("Hello, World!") //this will print Hello, World! in console
	// nm:="Parv"
	// fmt.Println("Hello,", nm) //this will print Hello, Parv
	// a,b:=5,6
	// fmt.Println(a+b,(a+b)/2) //expression as arg inside Println
	// //2 Printf :- used to print formatted output and doesnt add new line at the end just like Println
	// fmt.Printf("Hello, %s\n", nm) //this will print Hello, Parv
	// fmt.Printf("Sum: %d, Average: %d\n", a+b, (a+b)/2)
	// //fmt.Printf("xsknk """)//double quotes inside double quotes not allowed so solution is to use backslash before double quotes to escape it
	// fmt.Printf("xsknk \"\"") //this will print xsknk ""
	// //for quoted string we can use %q
	// fmt.Printf("Quoted string: %q\n", "Hello, World!") //this will print "Hello, World!"
	// // %v - can be replaced by any type
	// d,e,f:=1,"Fredwe",6.43
	// fmt.Printf("Values: %v, %v, %v\n", d, e, f) //this will print Values: 1, Fredwe, 6.43
	// // %T - will print the type of the value
	// fmt.Printf("Types: %T, %T, %T\n", d, e, f) //this will print Types: int, string, float64
	// // %t - will print the boolean value
	// fmt.Printf("Boolean: %t\n", true) //this will print Boolean: true
	// // %f - will print the float value
	// fmt.Printf("Float: %f\n", f) //this will print Float: 6.430000
	// // %s - will print the string value
	// fmt.Printf("String: %s\n", e) //this will print String: Fredwe	
	// //To convert base 10 to base 2 %b
	// fmt.Printf("Binary: %b\n", 10) //this will print Binary: 1010
	// //to print base 2 no with specific bits
	// fmt.Println("8 bit binary: %08b\n", 10) //this will print 00001010
	// //to print base 8 no
	// fmt.Printf("Octal: %o\n", 10) //this will print Octal: 12
	// //to print base 16 no
	// fmt.Printf("Hexadecimal: %x\n", 10) //this will print Hexadecimal: a
	// //to print base 16 no with capital letters
	// fmt.Printf("Hexadecimal: %X\n", 10) //this will print Hexadecimal: A
	// //to print base 16 no with prefix 0x
	// fmt.Printf("Hexadecimal with prefix: %#x\n", 10) //this will print Hexadecimal with prefix: 0xa
	// //to print base 16 no with prefix 0X
	// fmt.Printf("Hexadecimal with prefix: %#X\n", 10) //this will print Hexadecimal with prefix: 0XA
	// //to print base 16 no with prefix 0x and with specific bits
	// fmt.Printf("Hexadecimal with prefix and specific bits: %08x\n", 10) //this will print Hexadecimal with prefix and specific bits: 0000000a
	// //to print specfic no in float 
	// b:=3.28879879821897
	// fmt.Printf("Float with 2 decimal places: %.2f\n", b) //this will print Float with 2 decimal places: 3.29
	// //to print specfic no in float with 2 decimal places and with 0 padding
	// fmt.Printf("Float with 2 decimal places and 0 padding: %06.2f\n", b) //this will print Float with 2 decimal places and 0 padding: 03.29
	// //to print specfic no in float with 2 decimal places and with 0 padding and with prefix 0x
	// fmt.Printf("Float with 2 decimal places and 0 padding and with prefix: %#06.2f\n", b) //this will print Float with 2 decimal places and 0 padding and with prefix: 0x03.29
	
	/*Constants :- use this term to represent fixed(unchanging) val hence use const to avoid possible errors 
	All basic literals(1,3.4,"Hello") are unnamed constants
	A constant belongs to compile time and created at compile time and its val cant change while program is running 
	Constants need to initialized at the time of declaration
	in a grp constant if val of one const not defiend then it will take val of prev constant in the grp
	*/
	// const (
	// 	x=100
	// 	y=200
	// 	z
	// )
	// fmt.Println("Constants:", x, y, z) //this will print Constants: 100 200 200

	/*
	Constant Rules:-
	1. cant change val of const
	2. cant be initialise at runtime 
	3. cant use var to initialize const
	4. len function with fixed length of string can be initialized with const //als len is built-in function in go
	*/

	// //untyped const
	// const a int= 10 //typed constant
	// const b = 20 //untyped constant
	// fmt.Println("Constants:", a, b) //this will print Constants: 10 20
	// // //golang is strongly typed lang so we cant change type of const  ex:- 
	// // const c=3.2*a//this will throw compile time error as we are trying to assign float64 to int but 
	// // fmt.Println("Constant c:", c) //this will print Constant c: 32
	// //similarly
	// // const c float64=a//not permitted 
	// // fmt.Println(c)
	// const c float64=b//permitted
	// fmt.Println(c)
	// const d=3.2*b//this will not throw any error as we are assigning untyped constant to typed constant
	// fmt.Println("Constant d:", d) //this will print Constant d: 64
	
	/*Iota :- predefined identifier represents successive untyped int constants
	Its val is index of respective ConstSpec in that const declaration starting at 0,
	Can be used to construct a set of related constants*/
	// //example
	// const (
	// c1=iota
	// c2=iota
	// c3=iota
	// )
	// fmt.Println(c1,c2,c3) //Prints 0,1,2
	// const (
	// 	c5=iota
	// 	c6
	// 	c7
	// )
	// fmt.Println(c5,c6,c7)
	// //to get iota with 2 step
	// const (
	// 	c5=iota*2
	// 	c6
	// 	c7
	// )
	// fmt.Println(c5,c6,c7)
	// //to skip some val
	// const (
	// 	c5=iota*2
	// 	_ //to skip one val 
	// 	c6
	// 	c7
	// )
	// fmt.Println(c5,c6,c7)

	//Data Types :-  golang doesnt have char data type 
	/*1 Predeclared 
	int8,16,32,64 for signed int ,uint8,16,32,64 for unsigned int
	float32,64 : 0 before decimal pt can be ommitted
	complex64,128 
	byte : alias for unit8 
	rune :alias for int32
	bool :true or false 
	string :enclosed by ""
	*/
	// //Ex
	// var r rune ='f'
	// fmt.Printf("%T",r) //Print it as int32
	// fmt.Println(r) //Print dec ASCII value code for 'f'

	/*2 Composite Types
	array/slice :- numbered seq of elements of single type,array has fixed length but slice has dynamic length
	len is builtin funct to find length of an array
	map : unordered grp of eleements of one type,indexed by set of unique keys of another type,mP IN GO SIMILr to dict in python
	*/
	////Ex
	// var r=[4]int{1,2,-3,4}//array
	// fmt.Println(r)
	// var t=[]int{1,2,3,4,6}//slice
	// fmt.Println(t)

	// bal:=map[string]float64{ //map example
	// 	"d":2.2,
	// 	"e":3.2, //NOTE: , is necessary even after the last val
	// }
	// fmt.Println(bal)

	/*User DEfined type
	Struct Type : seq od named elements,called fields each of which has name and type,
	can be compared to class object in OOP
	Ex:	*/
	// type Car struct{
	// 	brand string
	// 	price int
	// }
	// var t Car
	// t.brand="wd"
	// t.price=32
	// fmt.Println(t)	

	/*
	Pointer data type : ptr is var that stores memory add of another var 
	val of an uninitalised ptr is nil
	Unlike C go has no ptr arithematic */
	// var x int =2
	// ptr :=&x
	// fmt.Printf("The type is %T and the val is %v as address real val will be %d",ptr, ptr,*ptr)

	//Function data type 
	//fmt.Printf("%T",ff)
/*Interfaces and function type
Channel type : provides mechanism for concurrently executing functions to communicate by sending and receiving values of specied element type*/

////GO Operators :-
// a:=11
// b:=14
// // operator :- symbol of prog lang which able to operate on values 
// // Types :- Arithmetic & Bitwise,Assignment,Inc & Dec statements,Comparison,Logical,pointer(&) & channels(<-)
// // Arithmetic :-
// fmt.Println(a+b)//addition
// fmt.Println(b-a)//subtraction
// fmt.Println(a*b)//multiplication
// fmt.Println(a/b)//division
// fmt.Println(a%b)//modulus or mod only apply on int 

// //Assignment operators :- 
// fmt.Println(a,b); //original values
// a+=2 //increement assignment
// fmt.Println(a,b);
// b-=1 //decreement assignment
// fmt.Println(a,b);
// a*=3 //multiplication assignment
// fmt.Println(a,b);
// b/=2 //division assignment
// fmt.Println(a,b);
// a%=b //mod assignment :- only remainder will be returned
// fmt.Println(a,b);
// /**Inc and dec statements :-  in golang they're statements not var like in cpp/java/python**/
// //fmt.Println(1+a--) //will print an error as ++ is statement not operator in golang
// a++
// b--
// fmt.Println(a,b);

// //Comparison operators: compare 2 operands and return bool val
// fmt.Println(a==b)	//equal
// fmt.Println(a>b)	//greater
// fmt.Println(a<b)	//smaller
// fmt.Println(a!=b)	//not equal
// fmt.Println(a>=b)	//greater or equal
// fmt.Println(a<=b)	//smaller or equal

// //Logical Operators :- apply to bool val and return bool vals
// fmt.Println(a>1 && b<0) //logical and 
// fmt.Println(a>1 || b<0) //logical or
// fmt.Println(!(a>1 || b<0)) //negate or not

// /*Overflow and underflow :-  if the bits of result goes beyond the resultant data type then higher bits or overflown bits values are discarded like in c/cpp
// Condition known as overflowing*/
// var x uint8=255 //max val of uint8
// x++
// fmt.Println(x)//x returns 0 val as its overflow and it'll not give runtime panick

// //go only catches error of overflow if the var is compiled at compile time 
// //Ex
// x:=int8(255)
// x++
// fmt.Println(x)//will catch the overflow error in compile time only
// f:=float32(math.MaxFloat32)
// fmt.Println(f)//will not give any error and overflows to infinity example :- 
// f*=2.35
// fmt.Println(f)//will print +inf 

// const i int8 =300 //it'll throw compilen time error 

////Converting Types
converting_types()

}
// func ff(){
// 	//for understanding function data type
// 	}

func converting_types(){
// //in golang we call casting type to converting type 
// x:=3 //int
// y:=4.33 //float
// //x*=y //throw compile type error due to golang being strongly typed lan which not allow int to * directly with float 
// x*=int(y)
// fmt.Println(x,y)
// fmt.Printf("%T %T\n",x,y)
// y=y*float64(x)
// fmt.Println(x,y)
// fmt.Printf("%T %T",x,y)

// //We cant assign var of different types except aliases type 
// var a =5 //int type
// fmt.Printf("%T\n",a)
// var b int64=6
// //a=b //error as both int and int64 is different even with same max and min
// a=int(b)//solution of the problem above
// fmt.Println(a,b)

//// Converting no to string and string to no 


}