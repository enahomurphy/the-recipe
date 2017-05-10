package helpers

import "fmt"

// Multiplier multiplies all
// integer parameters passed
// into this function
func Multiplier(args ...int) int {
	ans := 0
	for _, value := range args {
		ans += value
	}
	return ans
}

// PrintErr handles unexpected
// errors that occurs
func PrintErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
