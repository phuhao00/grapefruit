package testdata

import (
	"fmt"
	"strings"
	"testing"
)

//func removeChar(s string, index int) string {
//	if index < 0 || index >= len(s) {
//		return s
//	}
//	return s[:index] + s[index+1:]
//}

func TestStrRemove(t *testing.T) {
	str := "example"
	newIndex := 3  // 要插入字符的索引
	newChar := 'x' // 要插入的字符
	result := insertChar(str, newIndex, byte(newChar))
	fmt.Println(result) // 输出: "exaxmple"
}

func removeChar(s string, index int) string {
	if index < 0 || index >= len(s) {
		return s
	}

	var builder strings.Builder
	builder.WriteString(s[:index])
	builder.WriteString(s[index+1:])

	return builder.String()
}

func insertChar(s string, index int, char byte) string {
	if index < 0 || index > len(s) {
		return s
	}

	var builder strings.Builder
	builder.WriteString(s[:index])
	builder.WriteByte(char)
	builder.WriteString(s[index:])

	return builder.String()
}
