package main

import(
	"math"
	"fmt"
	"os"
)

func main(){

 //initialization of variables
	var a float64	//first number
	var b float64	//second number
	var c float64	//third number


	var d float64	//working result variable
	var e float64 = 0-b	//holds value for -b

	var f float64	//value of the result from addition
	var g float64	//value of the result from subtraction


		//variable inputs
	fmt.Print("Please enter your value for a:   ")
	fmt.Scan(&a)

	fmt.Print("\nPlease enter your value for b:   ")
	fmt.Scan(&b)

	fmt.Print("\nPlease enter your value for c:   ")
	fmt.Scan(&c)

	fmt.Print("\n")



		//checks if a is 0
	if(a == 0){
		fmt.Println("ERROR: the value for a cannot be 0")
		os.Exit(1)
	}

	d = (b*b) - (4*a*c)


		//checks if a sqrrt of a negative value will be taken
	if(d<0){
		fmt.Println("ERROR: cannot perform the squareroot of a negative number")
		os.Exit(1)
	}

	d =  math.Sqrt(d)

	f = e + d
	f = f/(2*a)

	g = e - d
	g = g/(2*a)


	fmt.Println("The results are:  ", f, " and ", g)


}
