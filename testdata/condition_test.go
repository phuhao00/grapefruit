package testdata

//
//import (
//	"fmt"
//	"testing"
//)
//
//func TestIf(t *testing.T) {
//	if true {
//		//todo
//	} else {
//		// 条件为假时执行的代码
//	}
//}
//
//func TestFor(t *testing.T) {
//	//for {
//	//
//	//}
//
//	for i := 0; i < 5; i++ {
//		fmt.Println(i)
//	}
//
//	var sliceDemo = make([]int, 0, 6)
//	for index, val := range sliceDemo { //有序
//		fmt.Println(index, val)
//	}
//	//
//	var mapDemo = make(map[string]int)
//
//	for key, val := range mapDemo {
//		fmt.Println(key, val)
//	}
//	var name = ""
//	switch name {
//	case "a":
//	case "b":
//	case "c":
//	default:
//
//	}
//
//}
//
//func TestGoto(t *testing.T) {
//	i := 0
//Loop:
//	for i < 5 {
//		fmt.Println(i)
//		i++
//		goto Loop // 无条件跳转到 Loop 标签处
//	}
//}
//
//func TestBreak(t *testing.T) {
//	for i := 0; i < 10; i++ {
//		fmt.Println(i)
//		if i == 5 {
//			break // 当 i 等于 5 时跳出循环
//		}
//	}
//}
//
//func TestContinue(t *testing.T) {
//	for i := 0; i < 5; i++ {
//		if i == 2 {
//			continue // 当 i 等于 2 时跳过当前迭代，开始下一次迭代
//		}
//		fmt.Println(i)
//	}
//}
//
//func TestFallthrough(t *testing.T) {
//	switch num := 2; num {
//	case 1:
//		fmt.Println("One")
//	case 2:
//		fmt.Println("Two")
//		fallthrough // 继续执行下一个 case
//	case 3:
//		fmt.Println("Three")
//	}
//}
