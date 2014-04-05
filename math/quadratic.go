package main

import (
	"fmt"
	"math"
)

// 二次项系数，一次项系数，常数项
type coefficient struct {
	first, second, third float64
}

// 求根的判别式的值
func (c *coefficient) detla () float64  {
	return math.Pow(c.second, 2) - 4 * c.first * c.third
}

// 判断方程是否有实根
func haveRoot(detla float64 ) bool {
	if detla >= 0{
		return true
	}

	return false
}

// 求方程的根
func (c *coefficient) root() []float64 {
	d := c.detla()
	r1 := (-c.second - math.Sqrt(d)) / (2 * c.first)
	r2 := (-c.second + math.Sqrt(d)) / (2 * c.first)
	return []float64{r1, r2}
}

func main() {
	co := coefficient{1.0, 5.0, 1.0}
	d := co.detla()
	fmt.Println("判别式为：", d)
	
	if haveRoot(d) {
		root := co.root()
		fmt.Println("方程的根为：", root)
	} else {
		fmt.Println("方程无实数根！")
	}
}
