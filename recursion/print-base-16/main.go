package main

import "fmt"

func main() {
	PrintBase16(16)
	fmt.Println()
	PrintBase16(160)
	fmt.Println()
	PrintBase16(17)
	fmt.Println()
}

func PrintBase16(n int) string {
	syms := "0123456789ABCDEF"
	base := 16
	r := n % base
	n = n / base

	if n != 0 {
		PrintBase16(n)
	}

	// select the appropriate symbol from syms, converting to string is needed since slice operator on a string returns rune
	fmt.Print(string(syms[r]))

	return ""
}

// The % operator gives us the remainder after dividing by our base, i.e. the leftmost digit
// We stpre the leftmost digit then repeat for the higher digits, this is accomplished through integer division by the base
// The clever part is having the recursive call before the print, this way the recursive call prints before the current code
// For example in base 10
// 987 % 10 = 7 <- we print this, the rightmost digit since we know there are 9 lots of 10 with a remainder of 8
// 987 / 10 = 98
// Recursive call
// 98 % 10 = 8
// 98 / 10 = 9
// Recursive call
// 9 % 10 = 9
// 9 / 10 = 0
// Recursive call exits
// Print(9)
// Print(8)
// Print(7)
