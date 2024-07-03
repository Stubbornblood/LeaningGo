GoSyllabus

var :
    can be used inside and outside of the function syntax var name type
    variable declare and value assign can be done seperately

:= :
    can only be used inside function
    variable declare and value assign cannot be done seperately only need to be done in same line


Use With printf
%v is used to print the value of the arguments
%T is used to print the type of the arguments

Go Slices:
    Similar to array but more powerfull and flexible
    Like array can store multiple values of same datatype in same variable
    Slices length can grow and shrink

    Ways to create Slices: 
        val = []type{values}
        Create Slices from array
        Using Make() function

Append to Slice:
    sliceArray := []int{1,2,3}
    sliceArray := append(sliceArray,4,5,6)
    output : [1 2 3 4 5 6]
     
    We can append two slice with each other also
    sliceArray1 = []int{1,2,3}
    sliceArray2 = []int{4,5,6}
    sliceArray3 = append(sliceArray1,sliceArray2...)
    output = [1 2 3 4 5 6]
    The '...' after slice2 is necessary when appending the elements of one slice to another.

Memory Efficiency:
    When working with slice Go loads all underlying elements into the Memory
    if array is large and you need only few elements, It is better to copy those elements using copy()
    The copy() function creates only new underlying array with only the required elements for the slice. This will reduce the Memory used for the program
        
Syntax To Rememeber:
    var (
		sum1 = 10+20
		sum2 = sum1 + sum1
		sum3 = sum2 + sum2
	)
	fmt.Print(sum3)

Loops:
	For loop is only available in GO

    for idx,value := range fruits{
        fmt.println(idx," ",val)
    }

Function:
    Note: When a parameter is passed to the function, it is called an argument. So, from the example above: fname is a parameter, while Liam, Jenny and Anja are arguments.