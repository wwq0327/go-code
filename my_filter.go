package main
import "fmt"

type testInt func(int) bool // 声明一个函数类型

func isOdd(integer int) bool {
	if integer % 2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer % 2 == 0 {
		return true
	}
	return false
}

func filter(f testInt, slice []int) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice = ", slice)
	odd := filter(isOdd, slice)
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(isEven, slice)
	fmt.Println("Even elements of slice are: ", even)
}
