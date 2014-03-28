package main

import "fmt"
import "math"

// p为用户自定义的比较精度， 比如 0.00001
func IsEqual(f1, f2, p float64) bool {
	return math.Abs(f1 - f2) < p
}

//  函数体只对数组进行复制，而不能修改原数组内容
func modify(arr [5]int) {
	arr[0] = 10
	fmt.Println("In modify(), arr values:", arr)
}


func main() {

	//bool 类型
	var v1 bool
	v1 = true
	v2 := (1 == 2) // 被推导为bool类型
	fmt.Println(v1, v2)

	//整型
	var value2 int32  //只初始化没被赋值的，默认值为0
	value1 := 64 //被自动推导为int类型

	fmt.Println(value2, value1)
	value2 = int32(value1) //强制类型转换
	fmt.Println(value2)

	// 数值运算： + - * / %
	fmt.Println(5 % 3)  //结果为2

	// 比较运算 > < == >= <= !=
	i, j := 1, 2
	if i == j {
		fmt.Println("i and j are equal")
	} else {
		fmt.Println("i and j are not equal")
	}

	//浮点数比较
	p := 0.000001
	var f1, f2 float64;
	f1 = 64
	f2 = 64.00001
	fmt.Println(IsEqual(f1, f2, p))

	// 字符串 
	var str string
	str = "Hello world"
	ch := str[0]
	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
	fmt.Printf("The First character of \"%s\" is %c.\n", str, ch)

	str1 := "hello, 世界"
	n := len(str)
	for i := 0; i<n; i++ {
		ch := str1[i] // 依据下标取字符串中的字符， 类型为byte
		fmt.Println(i, ch)
	}
	
	// 数组演示
	arr := [5]int{1, 2, 3, 4, 5}
	modify(arr)
	fmt.Println("In main(), arr values:", arr)

	// 数组切片
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 基于数组创建一个数组切片
	var mySlice []int = myArray[:5]
	fmt.Println("Elements of myArray:")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of mySlice:")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// 添加元素
	mySlice2 := []int{8, 9, 10, 11}
	mySlice = append(mySlice, mySlice2...)
	fmt.Println(mySlice, mySlice2)

	newSlice := mySlice2[:] // 复制所有元素
	fmt.Println(newSlice)

	// 内容复制
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	copy(slice2, slice1)
	copy(slice1, slice2)
	fmt.Println(slice2, slice1)
}
