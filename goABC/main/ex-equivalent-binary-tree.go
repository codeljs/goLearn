package main

import "fmt"

//import "golang.org/x/tour/tree"
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *Tree, ch chan int){
	if t == nil {
		return
	}
	if t.Left != nil{
		Walk(t.Left,ch)
		
	}
	ch <- t.Value
	if t.Right != nil{
		Walk(t.Right,ch)
	}


}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *Tree) bool{
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1,c1)
	go Walk(t2,c2)
	// a1 :=  make([]int, 10)
	// a2 :=  make([]int, 10)
	for i := 0; i < 10; i++ {
		a1 := <- c1
		a2 := <- c2 
		if a1 !=a2 {return false}
	}
	return true
}
// func Copy(array []int, ch chan int){

// }

func main() {
	fmt.Println()
}