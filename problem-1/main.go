package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str1, str2, op1, op2 string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nInput str1:")
	str1, _ = reader.ReadString('\n')

	fmt.Println("\nInput str2:")
	str2, _ = reader.ReadString('\n')

	op1 = compareStr(str1, str2)
	op2 = compareStr(str2, str1)
	fmt.Printf("\nop1: %v \nop2: %v\n\n", op1, op2)
}

func compareStr(a, b string) (ans string) {
	var flag bool
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if string(a[i]) == string(b[j]) {
				flag = true
				break
			} else {
				flag = false
			}
		}
		if !flag {
			ans = ans + string(a[i])
		}
	}
	return ans
}
