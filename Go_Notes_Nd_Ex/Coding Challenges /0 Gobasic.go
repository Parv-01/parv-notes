package main

import (
	"fmt"
	//"math"
	//"strconv"
)

/*for Ex4 */
//var version = "3.1" //will not provide error as error not allowed on package scope

func main() {
	fmt.Println("First Go challenger solutions")
	/*Ex1 */
	// var a int
	// var b float64
	// var c bool
	// var d string
	// x,y,z := 20,15.5,"Gopher!"
	// fmt.Println(a,b,c,d,x,y,z)

	/*Ex2 */
	// var (
	// 	a int
	// 	b float64
	// 	c bool
	// 	d string
	// )//right way
	// // var a,b,c,d int //wrong doing for multiple var declaration with single var
	// x,y,z:= 1,1,1
	// _,_,_,_,_,_,_=a,b,c,d,x,y,z

	/*Ex3*/
	// var a float64=7.1
	// x,y:=true,3.7
	// a,x=5.5,false //no new var on left side hence correct will be we changing val not again assigning them
	// fmt.Println(a,x,y)

	/* Ex4 */
	// name := "Golang"
	// fmt.Println(name )

	/*Ex5*/
	// const daysWeek,lightspeed,pi=7,299792458,3.14159
	// fmt.Println(daysWeek,lightspeed,pi)

	/*Ex6*/
	// const secinday,daysyr=31536000,365
	// fmt.Println(secinday*daysyr)

	/*Ex7*/
	//There will be only bool,int,float,rune,complex,string const
	//Slice cant be a constant

	/*Ex8*/
	// const(
	// 	jun=iota+6
	// 	july
	// 	aug
	// )
	// fmt.Println(jun,july,aug)

	/*Ex9*/
	// x,y,z:=10,15.5,"Gophers"
	// score:=[]int{10,20,30}
	// fmt.Printf("%d %f %s %v \n",x,y,z,score)
	// fmt.Printf("%q \n",z)//remember %q represent string ""
	// fmt.Printf("%v %v %v %v\n",x,y,z,score)
	// fmt.Printf("%T %T\n",y,score)

	/*Ex10*/
	// const l float64=1.422349587101
	// fmt.Printf("%0.4f\n",l)

	/*Ex11*/
	// var i int =3
	// var f float64 =3.2
	// f1:=float64(i)
	// f2:=int(f)
	// fmt.Printf("%T %T %f %d",f1,f2,f1,f2)

	/*Ex12*/
	// var i,f,s1,s2=3,3.2,"3.14","5"
	// fmt.Println(strconv.Itoa(i))
	// fmt.Println(strconv.Atoi(s2))
	// fmt.Println(strconv.FormatFloat(f,'f',1,64))
	// fmt.Println(strconv.ParseFloat(s1,64))

	/*Ex13*/
	// x,y:=4,5.1
	// z:=float64(x)*y
	// fmt.Println(z)
	// a:=5
	// b:=6.2*float64(a)
	// fmt.Println(b)

	/*Ex14*/
	// dist,spd:=149.6*math.Pow10(9),299792458
	// time :=float64(dist)/float64(spd)
	// fmt.Println(time)
	// //for better reading we can use _ between no example ;- 455_555_555

	/*Ex15 Good question from operators and conversions.html exercise 5 (48)*/
	// x,y:=0.1,5
	// var z float64
	// r1:=x<z || int(x)!=int(z)
	// r2:=y==1*5 && int(z)==0
	// fmt.Println(r1,r2)

	/*Ex16*/
	// type dur int
	// var hr dur
	// fmt.Println(hr)
	// fmt.Printf("%T \n",hr)
	// hr=3600
	// fmt.Printf("%v %T",hr,hr)

	/*Ex17*/
	// type (
	// 	km float64
	// 	ml float64
	// )
	// const m2km=1.609
	// var mbtp ml =655.3
	// var kmbtp km = km(mbtp)*km(m2km)
	// fmt.Println(kmbtp)

}
