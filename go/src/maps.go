package main

import(
	"fmt"
)

func main(){
	// lookup:=make(map[string]int)
	// lookup["goku"] = 9001

	lookup := map[string]int{
		"goku": 9001,
		"gohan": 2044,
	}

//below is to iterate over a map to return key value pair in random order
// for key, value := range lookup{
//	...
//}


	power, exists := lookup["vegeta"]

	// for number of keys
	// total := len(lookup)

	//to delete
	// delete(lookup, "goku")

	fmt.Println(power, exists)
}


//to use map as a field in a structure:
// 	type Saiyan struct {
// 		Name string
// 		Friends map[string]*Saiyan
// 	}

	