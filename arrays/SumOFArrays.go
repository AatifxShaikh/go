package main

func Sum(numbers []int) int {
	var result int
	for _, number := range numbers {
		result += number
	}
	return result
}
func SumAll(array1 []int, array2 []int) []int {
	var sum []int
	val := Sum(array1)
	val2 := Sum(array2)
	sum = append(sum, val)
	sum = append(sum, val2)
	return sum
}
func SumOFTail(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
