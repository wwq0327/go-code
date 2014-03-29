package main
import "fmt"

func main() {
	s := "hello"
	c := []byte(s) // 将string 转换成 []byte
	fmt.Println(c, s)
	c[2] = 'c'
	s2 := string(c) // 将[]byte转换成string
	fmt.Println(s2)

	// map 声明

	m1 := map[string]int {"C": 2, "Go": 5}
	fmt.Println(m1)

	m2 := make(map[string]int) // 先声明，再使用，使用make()
	m2["1"] = 1
	fmt.Println(m2)
}
