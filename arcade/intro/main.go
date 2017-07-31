package intro

import "math"

//01
func add(param1 int, param2 int) int {
	return param1 + param2
}

//02
func centuryFromYear(year int) int {
	if math.Mod(float64(year), 100) == 0 {
		year = year / 100
	} else {
		year = (year / 100) + 1
	}
	return year
}

//03
func checkPalindrome(inputString string) bool {

	n := 0
	rune := make([]rune, len(inputString))
	for _, r := range inputString {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	output := string(rune)
	if output == inputString {
		return true
	}
	return false
}

//04
func adjacentElementsProduct(inputArray []int) int {

	largest := inputArray[0] * inputArray[1]
	tmp := 0

	for i := range inputArray {
		if i+1 < len(inputArray) {
			tmp = inputArray[i] * inputArray[i+1]
			if tmp > largest {
				largest = tmp
			}
		}
	}
	return largest
}

//05
func shapeArea(n int) int {
	return (n * n) + ((n - 1) * (n - 1))
}
