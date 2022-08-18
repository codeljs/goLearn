package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func foo0() func() {
    x := 1
    f := func() {
        fmt.Printf("foo0 val = %d\n", x)
    }
    x = 11
    return f
}

func fibonacci() func() int {
	x1 := 0
	x2 := 1
	f := func() int{
		res := x1+x2
		x1 = x2
		x2 = res
		return res
	}
	return f
}


func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	foo0()() // 延迟绑定 late binding
}