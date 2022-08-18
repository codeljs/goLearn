package main

import "fmt"

func main(){
	D := make([]int, 3, 4)
	E := make([]int, len(D))
	for i := 0; i < len(D); i++ {
		D[i] = i
	}
	copy(E, D)
	fmt.Printf("D = %v, len = %d, cap = %d\n", D, len(D), cap(D))

	E = append(E, 3)
	fmt.Printf("D: %v\n", D)
	fmt.Printf("E: %v\n", E)

	fmt.Println("Setting a different value for E...")
	E[0] = 4
	fmt.Printf("D: %v\n", D)
	fmt.Printf("E: %v\n", E)

	fmt.Printf("D: %p\n", D)
	fmt.Printf("E: %p\n", E)
	fmt.Println()
	for idx := range D {
		fmt.Printf("&D[%d] = %p\n", idx, &D[idx])
	}
	fmt.Println()
	for idx := range E {
		fmt.Printf("&E[%d] = %p\n", idx, &E[idx])
	}



}