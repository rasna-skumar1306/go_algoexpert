package main

//IsValidSubsequence is a function that checks if the subsequence is present in main array also is it a valid one
func IsValidSubsequence(array []int, sequence []int) bool {
	seqID := 0
	for _, num := range array {
		if seqID == len(sequence) {
			return true
		}
		if sequence[seqID] == num {
			seqID++
		}
	}
	return seqID == len(sequence)
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10} //sample input array
	sequence := []int{1, 6, -1, 10}            //sample sequence
	print(IsValidSubsequence(array, sequence))
}
