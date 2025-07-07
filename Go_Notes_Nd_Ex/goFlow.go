package main

import (
	"fmt"
	//"os"
	"strconv"
)

func main(){
	fmt.Println("Ich leibst koding.")
	//If Else and else if conditions
	//ifElse()

	//Command Line arguments{User input} use os.Args
	//cmdlinearg()

	//Simple If or short way of writing if statements
	// simpleIf()

	//For Loop and loops
	forLoop()
}

// func ifElse(){
// 	/*if some_cond{}
// 	else if some_con{}
// 	else {}*/
// 	price,stock:=100,true
// 	if price>120{
// 		fmt.Println("its there")
// 	}else if price==110{
// 		fmt.Println("itsnot here")
// 	}else{
// 		fmt.Println(stock)
// 	}
// 	if price ==120 || stock{
// 		fmt.Println("one is there")
// 	}
// 	if price ==100 && stock{
// 		fmt.Println("BOTH THERE")
// 	}
// }

// func cmdlinearg(){
// 	/*To see the value of os.Args
// 	When using through cmdline when written go run goFlow.go x y z then
// 	op will be os.Args val and then all arguments provided afterwards */
// 	//fmt.Println("os.Args ",os.Args)
// 	//
// 	/*When run in cmdline :- go run goFlow.go abcdb abhbdcjsb cdcjknjk
// 	Path:  /home/pro/.cache/go-build/38/3854e917565a14a46fa039ffe53e043dcb66355836f79bc3fd97dafd5fd7c28e-d/goFlow
// 	1 arg:  abcdb
// 	2 arg:  abhbdcjsb
// 	No of items in os.Args:  4
// 	*/
// 	// fmt.Println("Path: ",os.Args[0])
// 	// fmt.Println("1 arg: ",os.Args[1])
// 	// fmt.Println("2 arg: ",os.Args[2])
// 	// fmt.Println("No of items in os.Args: ",len(os.Args))
// 	//To find the type of os.Args
// 	//fmt.Printf("%T",os.Args)
// }

// func simpleIf(){
// 	//Short statements for if in golang
// 	if i,err:= strconv.Atoi("44");err==nil{ //First statement is innitialisation statement for local var,2 part is boolean statement
// 		fmt.Println("No error i is ",i)
// 	}else{
// 		fmt.Println("ERROR!! ",err)
// 	}
//
// 	//This complete above code behaves same as :- 
// 	/* i,err := strconv.Atoi("44")
// 	if err!=nil{
// 	fmt.Println("NO error ",i)}
// 	else{
// 	fmt.Println("ERROR!!! ",err)}
// 	*/
// }

func forLoop(){

}