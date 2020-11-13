package main

import "fmt"


func main(){
	array1 := [5]string{"i","n","h","a","!" }
	slice1 := array1[0:3]
	
	fmt.Println(slice1)
	fmt.Println(array1)
	fmt.Println(array1[2:])
	fmt.Println(slice1[:2])

	array2 := [5]string{"a","b","c","d","e"}
	
	slice2 := array2[0:3]
	slice3 := array2[2:5]
	
	fmt.Println(slice2,slice3)
	
//	array2[2] = "z"
	slice2[2] = "King"
//	slice2 := array2[0:3]
//	slice3 := array2[2:5]

	fmt.Println(slice2,slice3)
	fmt.Println(array2)

}
