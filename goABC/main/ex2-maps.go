package main

import (
	"strings"
	//"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	aux := strings.Fields(s)
	dict := make(map[string]int)
	for i:=0;i<len(aux);i++{
		dict[aux[i]] ++
	}
	return dict
}

func main() {
	//wc.Test(WordCount)
}
