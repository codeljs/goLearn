package main

import "fmt"

func main() {
	

	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	temp := s

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	s[0] = 2
	printSlice(s)
	printSlice(temp)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var p []int = make([]int,5,100 )
	printSlice(p)
	p = primes[0:]
	fmt.Println(p)
	p = append(p,p... )
	printSlice(p)
	p = append(p,1)
	printSlice(p)
	fmt.Println(primes)
	


	
	D := make([]int, 3, 4)
	for i := 0; i < len(D); i++ {
		D[i] = i
	}
	fmt.Printf("D = %v, len = %d, cap = %d\n", D, len(D), cap(D))
	
	E := append(D, 3)
	fmt.Printf("D: %v\n", D)
	fmt.Printf("E: %v\n", E)
	
	fmt.Println("Setting a different value for E...")
	E[0] = 4
	fmt.Printf("D: %v\n", D)
	fmt.Printf("E: %v\n", E)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
