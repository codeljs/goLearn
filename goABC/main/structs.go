package main

import "fmt"

type Vertex struct {
	X int
	Y int
	Z string
}

func main() {
	temp := Vertex{3,3,"a"}
	temp.Z = "change"
	fmt.Println(temp)
	fmt.Println(Vertex{1, 2,"abc"})

}