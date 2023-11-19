package testdata

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMapDemo(t *testing.T) {
	runtime.GOMAXPROCS(4)
	var myMap = make(map[string]int, 6)
	// 初始化 map，存储键值对
	myMap["apple"] = 10
	myMap["banana"] = 5
	myMap["orange"] = 7
	myMap["orange1"] = 7
	myMap["orange2"] = 7
	myMap["orange3"] = 7
	myMap["orange4"] = 7
	delete(myMap, "apple")
	i, exists := myMap["apple"]
	if exists {
		fmt.Println(i)
	} else {
		fmt.Println("not exists")

	}
	for s, i2 := range myMap {
		fmt.Println(s, i2)
	}
}
