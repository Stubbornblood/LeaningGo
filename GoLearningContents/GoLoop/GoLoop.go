package main
import ("fmt")

func main(){
	// for i:=1; i<5; i++{
	// 	fmt.Print(i," ")
	// }
	// fmt.Print("\n")
	// for i:=10; i<100;i+=10{
	// 	fmt.Print(i," ")
	// }


	//Continue
	// for i:=0; i<10; i++{
	// 	if i == 3{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//break
	// for i:=0; i<5; i++{
	// 	if i == 3{
	// 		break
	// 	}
	// 	fmt.Print(i)
	// }

	// adj := [...]string{"big","small"}
	fruits := [...]string{"Apple","Mango","Banana"}

	// for i:=0; i<len(adj); i++{
	// 	for j := 0 ; j<len(fruits); j++{
	// 		fmt.Println(adj[i],fruits[j])
	// 	}
	// }

	// for i,d := range fruits{
	// 	fmt.Println(i," ",d)
	// }

	// for _,d := range fruits{
	// 	fmt.Println(d)
	// }

	for d, _ := range fruits{
		fmt.Println(d)
	}


}