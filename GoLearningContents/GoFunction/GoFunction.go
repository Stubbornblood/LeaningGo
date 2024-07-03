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


func main(){
	familyNames("Liam",23)	
	fmt.Println(Myfunction(10,20))
	fmt.Print(MySecondFunction(20,30))
}
