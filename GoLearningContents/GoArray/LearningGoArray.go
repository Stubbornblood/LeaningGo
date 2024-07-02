package main
import ("fmt")

func main(){

	var arr1 = [3]int{1,2,3}
	fmt.Printf("%v",arr1)
	arr2 := [3]int{1,2,3}
	fmt.Println(arr2)
	arr3 := [...]int{1,2,3,4}
	fmt.Println(arr3)
	fmt.Print(arr3[2])
	fmt.Print("\n")
	arr3[2] = 5
	fmt.Print(arr3[2])
	fmt.Print("\n")
	//if the array value is not initialize
	arr4 := [5]int{}
	fmt.Print(arr4)
	fmt.Print("\n")
	arr5 := [5]int{1,2,3}
	fmt.Print(arr5)

	//initializing only specific element
	arr6 := [5]int{1:10,2:20}
	fmt.Print(arr6)
	fmt.Print(len(arr1))
	fmt.Print("\n")
	arr7 := [...]int{}
	print(len(arr7))
	

}