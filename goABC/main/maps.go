package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	var n = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967}, //若顶级类型只是一个类型名，你可以在文法的元素中省略它。
		"Google":    {37.42202, -122.08408},
	}
	n["Bell Labs"] = Vertex{
		40, -74,
	}
	
	fmt.Println(m["Bell Labs"])
	fmt.Println(n)
}
