package main
import (
	"fmt"
)

func main() {
	y := 37
	pi := &y // 取y的地址，pi即是一个指向一个整型变量y的指针
	ppi := &pi // ppi 为一个指向一个指针类型 pi 的指针
	fmt.Println(y, pi, ppi)

	y += 1 
	fmt.Println(y, pi, ppi)
	
	*pi += 1 // 取pi所指向量的的值，并加1
	fmt.Println(y, pi, ppi)

	**ppi += 1 // 取ppi所指向的指针所指向的值，并加1
	fmt.Println(y, pi, ppi)

	type pointer struct {
		name string
	}

	ps := new(pointer) // 取pointer的地指址，ps为一个指向pointer的结构型指针
	fmt.Printf("type: %T, value: %p\n", ps, ps)
	
	// 自动解引用
	ps.name = "wyatt" // 等效于(*ps).name  显式解引用
	fmt.Println(*ps, ps) // 打印指向的值 {wyatt} &{wyatt}
}

