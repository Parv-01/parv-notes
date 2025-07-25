package codeChallenge
import(
	"fmt"
)
func ChallSlices(){
	/*EX1*/
	fmt.Println("Challenges for Slices")
	a:=[]string{"a","b","c"}
	fmt.Println(a)
	for i,v := range a{
		fmt.Println(i,v)
	}
	/*EX3*/
	nums:=[]float64{1.1,2.2,3.3}
	nums=append(nums,10.1)
	nums=append(nums,4.1,5.5,6.6)
	fmt.Println(nums)
	n:=[]float64{0.0,0.0}
	nums=append(nums,n...)
	fmt.Println(nums)
	/*EX4 HAve to work with it not completed yet*/
	tre:=[]int{}
	fmt.Scanln(tre)
	
}