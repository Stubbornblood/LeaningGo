package main
import ("fmt")

func main(){
	// SliceArray := []int{10,20,30,40}
	// fmt.Println(SliceArray[0:3])
	// fmt.Print(cap(SliceArray))
	// SliceArray[2] = 2323
	// fmt.Println(SliceArray)
	// fmt.Print(cap(SliceArray))
	// //Append Element to Slice
	// newSlice := append(SliceArray,12,13,14)
	// fmt.Print(newSlice)
	// fmt.Print(cap(newSlice))

	// SliceArray1 := []int{1,2,3}
	// SliceArray2 := []int{4,5,6}
	
	// SliceArray3 := append(SliceArray1,SliceArray2...)
	// fmt.Print(SliceArray3)


	//Change the length of the Slice


	arr1 := [6]int{9,10,32,44,55,66}
	SliceArray1 := arr1[1:3]
	fmt.Println(SliceArray1)
	fmt.Println(len(SliceArray1))
	fmt.Println(cap(SliceArray1))

	mySlice1 := append(SliceArray1,23,343,544,545)
	fmt.Println(mySlice1)
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))
}
