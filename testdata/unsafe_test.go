package testdata

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {

	str := "Hello, unsafe!"

	// 将字符串转换为字节切片（[]byte）
	byteslice := []byte(str)

	// 将字节切片转换为字符串（不安全操作）
	unsafeString := bytesToString(byteslice)

	fmt.Println("Original string:", str)
	fmt.Println("Unsafe conversion:", unsafeString)

	// 检查两个字符串是否相等
	fmt.Println("Are strings equal?", str == unsafeString)
	//json.Marshal()
	//json.Marshal()
}

func bytesToString(b []byte) string {
	// 使用 unsafe 将字节切片转换为字符串
	return *(*string)(unsafe.Pointer(&b))
}
