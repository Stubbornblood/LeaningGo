package main
import ("fmt")

func main(){

	//Copy
	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	// fmt.Print(numbers)
	// fmt.Print(len(numbers))
	// fmt.Print(cap(numbers))

	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int,len(neededNumbers))
	fmt.Println(neededNumbers," ",numbersCopy)

	copy(numbersCopy,neededNumbers)
	fmt.Println(numbersCopy)
	fmt.Println(len(numbersCopy))
	fmt.Println(cap(numbersCopy))
}