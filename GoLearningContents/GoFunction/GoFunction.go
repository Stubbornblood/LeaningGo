package main
import ("fmt")


// func myMessage(){
// 	fmt.Print("I Just got executed")
// }

func familyName(fname string){
	fmt.Println("Hello",fname,"Refsness")
}

func familyNames(fname string,age int){
	fmt.Println("Hello",fname,"My age is : ",age)
}

func Myfunction(a int,b int)int{
	return a+b
}

func MySecondFunction(a int,b int)(result int){
	result = a+b
	return
}

func MyThirdFunction(a int,b int)(result int){
	result = a+b
	return
}

func MyFourthFunction(a int,b int)(result int , txt1 string){
	result = a+b
	txt1 = "Hello"
	return
}

func main(){
	familyNames("Liam",23)	
	fmt.Println(Myfunction(10,20))
	fmt.Println(MySecondFunction(20,30))
	val := MyThirdFunction(40,50)
	fmt.Println(val)
	a,b := MyFourthFunction(100,2000)
	fmt.Println(a," ",b)
	_,b = MyFourthFunction(10,23)
	fmt.Println(b)
}
