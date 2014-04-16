package mymap

func UniqueInts(slice []int) []int {
	seen := make(map[int]bool)
	unique := make([]int, 0)
	for _, x := range slice {
		if _, found := seen[x]; !found {
			unique = append(unique, x)
			seen[x] = true
		}
	}
	return unique
}
