package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, 世界",rand.Int63n(10))
	fmt.Println(swap("1","2"))
	fmt.Println(add(1,2))

}
func add(x , y int) (z int){
	return x + y;
}
func f(func (x int, y int) int, int) func(int, int) int{
	var p *int
	*p = 3

	return add;

}
func swap(x, y string) (string, string) {  //mutilple return value
	return y, x
}