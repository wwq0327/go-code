package main
// 数组定义的两种方式

import "fmt"

func main() {
    var a [5]int
    fmt.Println(a)

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println(b)

    // 创建个长度为三的空数组
    s := make([]string,  3)
    fmt.Println(s, len(s))
    s = append(s, "apple")
    fmt.Println(s)

    s = append(s, "red")
    fmt.Println(s)

    s[0] = "hahaha"
    fmt.Println(s)
}