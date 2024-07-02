package main
import ("fmt")

func main(){
	arrSlices := []int{}
	fmt.Println(len(arrSlices))
	fmt.Println(cap(arrSlices))

	arrSlices2 := []string{"Go","Slices","Are","Powerful"}
	fmt.Println(arrSlices2)
	fmt.Println(len(arrSlices2))
	fmt.Println(cap(arrSlices2))

	arrSlices3 := []int{2:4}
	fmt.Print(arrSlices3)

	arr := [5]int{1,2,3,4,5}
	arrSlices4 := arr[2:5]
	fmt.Print(arrSlices4)
}