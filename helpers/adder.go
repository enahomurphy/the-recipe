package helpers

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
