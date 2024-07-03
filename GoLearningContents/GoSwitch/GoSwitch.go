package main
import ("fmt")

func main(){
	// day := 7

	// switch day{
	// case 1:
	// 	fmt.Print("Monday")
	// case 2:
	// 	fmt.Print("Tuesday")
	// case 3:
	// 	fmt.Print("Wednesday")
	// case 4:
	// 	fmt.Print("Friday")
	// case 5:
	// 	fmt.Print("Saturday")
	// case 6:
	// 	fmt.Print("Sunday")
	// default:
	// 	fmt.Print("Not a Weekday")
	// }

	// day := 7

	// switch day{
	// case 1:
	// 	fmt.Print("Monday")
	// case 2:
	// 	fmt.Print("Tuesday")
	// default:
	// 	fmt.Print("No week day available")
	// }


	day := [...]int{1,0}


	switch day{
	case [...]int{1,2}:
		fmt.Print("Hello World")
	case [...]int{5,4}:
		fmt.Print("No Hello")
	default :
		fmt.Print("Chank lan day")
	}





}