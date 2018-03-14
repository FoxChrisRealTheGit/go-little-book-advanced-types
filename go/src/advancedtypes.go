package main

// import (
// 	"fmt"
// )


// func main(){
// //arrays
// // var scores [10]int
// // score[0] = 338

// // scores:=[4]int{9001, 9333, 212, 33}


// //slice
// // scores := []int{1,4,293,4,9}

// //below is length of 10 and capacity of 10
// // scores:= make([]int, 10)

// //bellow is length of 0 and capacity of 10
// scores := make([]int, 0, 10)


// //below cause error
// // scores[7] = 9033

// //below will work but place in position 0
// // scores = append(scores, 5)

// scores = scores[0:8]
// scores[7] = 9033

// //four common ways to initialize a slice
// // names:=[]string{"Name", "Names", "MoarNames"}
// // checks := make([]bool, 10)
// // var names []string
// // scores := make([]int, 0, 20)

// 	fmt.Println(scores)
// }