package main
import ("fmt")

func main(){
	// var a = 10
	// fmt.Print(a)

	// var a = 10
	// a += 5
	// fmt.Print(a)

	//compare Operator
	//  var x = 5
	//  var y = 3
	//  fmt.Println(x>y)

	 //BitWise Operator
	 var x = 9
	 var y = 8
	 fmt.Printf("x = %b\n",x)
	 fmt.Printf("y = %b\n",y)
	 fmt.Printf("x & y is %b\n",x & y)
	 fmt.Printf("%b\n",x|y)
	 fmt.Printf("%b \n",x^y)
	 fmt.Printf("%b\n",x<<2)
	 fmt.Printf("%b",x>>2)
}