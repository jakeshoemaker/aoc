package helpers

import "fmt"

// joins two strings together
func JoinStr(one string, two string) string {
	str := fmt.Sprintf("%v%v", one, two)
	return str
}

// doubles whatever number you send
func Double(num int64) int64 {
	return num * 2
}
