package main

import "fmt"

func main() {
	//Declaration of a slice
	mySlices := []int{1, 2, 3, 4, 5}

	// want to append add to the slice
	mySlices = append(mySlices, 6, 4, 32, 2)
	fmt.Println(mySlices)

	// from the first to the 4 using 1st index
	cut1 := mySlices[:4]
	fmt.Println(cut1)

	// from the 4 using 0 index to the last
	cut2 := mySlices[4:]
	fmt.Println(cut2)

	//loopig through the slice
	// for i := 0; i < len(mySlices); i++ {
	// 	fmt.Printf("At the %v index the value is: %v \n", i, mySlices[i])
	// }

	//the len and cap
	fmt.Printf("Slice:%v The len of the slice:%v and the cap:%v\n", mySlices, len(mySlices), cap(mySlices))
	fmt.Printf("The slice:%v len:%v cap:%v \n", cut1, len(cut1), cap(cut1))
	fmt.Printf("The slice:%v len:%v cap:%v \n", cut2, len(cut2), cap(cut2))

}
