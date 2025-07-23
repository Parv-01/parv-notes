package main

import ( //FILE SCOPE
	"fmt"

	f "fmt"
)

//permitted and used to learn about scopes

const done = false //package type scope
var a = 5          //will not return error even if not used :- package scope

func GoFlow() {
	f.Println("Ich leibst koding.//Go flow")
	fmt.Println("Ich leibst koding.//Go flow")
	/*If Else and else if conditions*/
	//ifElse()

	/*Command Line arguments{User input} use os.Args*/
	//cmdlinearg()

	/*Simple If or short way of writing if statements*/
	// simpleIf()

	/*For Loop and loops*/
	// ForLoop()

	/*Continue  and Break Statement */
	// continueBreakstat()

	/*Label Statements */
	// labelstat()

	/*goto and switch statements*/
	//gotoSwitchstat()

	/*Scope in Golang*/
	// 	scopego()
	// 	var task = "running" //local scope /block scope :- remains visible until the {} of main()
	// 	f.Println(task,done)
	// 	const done =true//valid as scope is different :- local scoped
	// 	f.Println("done in main",done)
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

// func ForLoop(){
// 	//for loop
// 	for i :=0;i<10;i++{ //1 is init statement with temp var i which only exist inside for loop, 2 is test condition ,3 process statement
// 		fmt.Println(i)
// 	}
// 	//No while loop in go but we can construct that
// 	j:=10
// 	for j<14{
// 		fmt.Println(j)
// 		j++
// 	}
// }

// func continueBreakstat(){
// 	//continue statemenent example
// 	for i:=0;i<10;i++{
// 		if i%2==0{
// 			fmt.Println(i)
// 		}else{
// 			continue //will skip all the functions afterwards and next iteration of loop will start
// 			i++
// 		}
// 	}
// 	//
// 	// Break Statement Example
// 	for i:=0;i<10;i+=2{
// 		if i==8{
// 			fmt.Println("wohhoo break works!!!")
// 			break //not end the whole program just came out of process/loop
// 		}else{
// 			continue
// 		}
// 	}
// }

// func labelstat(){
// 	/*Label stat :-  used in break,continue,goto statements
// 	ILLEGAL to define label statements which is never used
// 	They are not block scoped and dont conflict with identifiers that are not labels ,They live in naother space
// 	Scope of label is body of fucntion in which its declared and excludes body of any nested function
// 	Most times labels are used to terminate outer enclosing loops
// 	*/
// 	outer :=19
// 	_=outer //this will not conflict with identifiers and label with same name hence no error will be seen
// 	people :=[5]string{"n","b","a","d","e"}
// 	frien := [2]string{"a","t"}
// outer://label tells compiler that we care about this place
// 	for index,name :=range people{//range used to iterate over arrays and return both index and elements one by one from array
// 		for _,frind := range frien{
// 			if name==frind{
// 				fmt.Printf("found friend at %q at %d \n",frind,index)
// 				break outer //without label it breaks from inner loop only
// 				//ith label it will directly comes to outer label statement
// 			}
// 		}
// 	}
// 	fmt.Println("Next instruction after the break")
// }

// func gotoSwitchstat(){
// 	/*GOto :- transfer control to statement with corresponding label within the same fucntion
// 	It allows us to jump to any label inside the function but break,continue restricted to if,else,for but goto is not */
// 	i:=0
// loop: //creates loop like for without using for
// 	if i<5{
// 		fmt.Println(i)
// 		i++
// 		goto loop
// 	}
// 	//NOTE :- its not allowed to jump to label if new var are introduced
// 	// goto todo
// 	// x:=5
// 	// todo:
// 	// 	fmt.Println("ewndjhdb")
// 	//
// 	//Goto shuld be used with caution
// 	//
// 	/*Switch statements :-  in backend golang converts in nested if only
// 	Golang adds break automatically aqafter each case*/
// 	lang:="golang"
// 	switch lang{
// 	case "py":
// 		fmt.Println("py")
// 	case "GO","Go","golang":
// 		fmt.Println("Goood!!!")
// 	default:
// 		fmt.Println("HAVE SOME LIFE BRO!!!!!")
// 	//Cases are evaluated from top to bottom and default is not necessary
// 	}
// 	//
// 	n:=5
// 	switch{ //ALways true kind of switch statement or can also be written as
// 		//switch true{}
// 	case n%2==0:
// 		fmt.Println("true")
// 	case n%2!=0:
// 		fmt.Println("odd")
// 	}
// }

// func scopego(){
// 	/*Scope :- means visibility,also called as lifetime  of var:- is interval of time during which it exists as program executes
// 	Name cant be declared again in same scope ,but can be declared in another scope
// 	3 types of scopes:-
// 	File,Package,Block(local)
// 	*/
// 	//So we cant import "fmt" again in file but alias can be done import f "fmt"
// 	f.Println("Here the value printed for done const will be of package scoped only{no local scope value WILL BE PRINTED HERE}",done)
// 	//Unused package scope var doesnt return error
// }
