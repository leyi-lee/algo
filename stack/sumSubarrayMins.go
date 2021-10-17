package stack

func sumSubarrayMins(arr []int) int {
	sum := 0
	s := []int{}
	arr = append(arr, 0)
	for i :=0; i < len(arr); i++ {
		for len(s) > 0 && arr[s[len(s)-1]] >= arr[i] {

		}
	}

	return sum
}