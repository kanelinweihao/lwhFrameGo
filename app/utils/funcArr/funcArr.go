package funcArr

// TrimZero trim()
func TrimZero(arr []string) []string {
	for len(arr) > 1 && arr[0] == "0" {
		arr = arr[1:]
	}
	for len(arr) > 1 && arr[len(arr)-1] == "0" {
		arr = arr[:len(arr)-1]
	}
	return arr
}

// ReverseSlice reverse()
func ReverseSlice(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
