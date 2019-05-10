package main

func main() {}

func arraySum(array []int) int {
	sum := 0
	for _, number := range array {
		sum += number
	}
	return sum
}

// MaximumSubarraySum ...
func MaximumSubarraySum(numbers []int) int {
	topSum := 0

	for idx1 := 0; idx1 < len(numbers); idx1++ {
		for idx2 := len(numbers); idx2 > idx1; idx2-- {

			newSum := arraySum(numbers[idx1:idx2])

			if newSum > topSum {
				topSum = newSum
			}
		}
	}
	return topSum
}
