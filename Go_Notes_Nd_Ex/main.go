package main

import (
	"github.com/Parv-01/parv-notes/Go_Notes_Nd_Ex/codeChallenge"//For importing all the files inside the Challenges Folder
	"fmt"
)

func main() {
	fmt.Println("Main file to run whole package.")
	/*Go Basics*/
	// GoBasics()
	/*Go Flow*/
	// GoFlow()
	/*Go Arrays*/
	// GoArrays()
	/*Go Slices*/
	// GoSlices()

	/*Challenge Exercises*/
	Challenges()
}
func Challenges(){
	/*Challenge 1*/
	// codeChallenge.ChallBasics()
	/*Challenge 2*/
	// codeChallenge.ChallFlow()
	/*Challenge 3*/
	// codeChallenge.ChallArrays()
	/*Challenge 4*/
	codeChallenge.ChallSlices()
}
