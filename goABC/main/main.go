package main

import ("fmt" 
		 "math/cmplx")

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
var c , python, java bool  = true, false,false//var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
func and( a bool, b bool)(bool){
	res:= true
	if(a&&b){
		res = true
		return res
	}
	res = false
	return res
}

func main() {
	fmt.Println(split(17))
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	fmt.Println(i,f,b,s)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

}
