package main
import ("fmt")

func main(){	
	var a = 15 + 12
	fmt.Println(a)

	var (
		sum1 = 10+20
		sum2 = sum1 + sum1
		sum3 = sum2 + sum2
	)
	fmt.Println(sum3)

	data := 12
	data --
	fmt.Print(data)
}