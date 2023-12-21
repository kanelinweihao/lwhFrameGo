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
