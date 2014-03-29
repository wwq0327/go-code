package main
import "fmt"

const(
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

// 自字一个Color类型，其实就是byte的别名
type Color byte

// 定义一个Box类型
type Box struct {
	width, height, depth float64
	color Color
}

type BoxList []Box // BoxList，一个slice，内容为Box类型

// 计算Box的体积
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

// 给Box指定颜色,需要修改box颜色，所以需要调用Box指针
// 当接收者为指针时，Go会自动以指针的方式去调用相应值
func (b *Box) SetColor(c Color) {
	b.color = c // 等效于 *b.color = c
}

// 体积最大的一个box的颜色
func (bl BoxList) BiggestsColor() Color {
	v := 0.0
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

// 给所有的box着色为BLACK
func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		// 以指针方式去修改值，等效于&bl[i].SetColor(BlACK)
		bl[i].SetColor(BLACK)
	}
}

// 根据值打印相应的颜色
func (c Color) String() string {
	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList {
        Box{4, 4, 4, RED},
        Box{10, 10, 1, YELLOW},
        Box{1, 1, 20, BLACK},
        Box{10, 10, 1, BLUE},
        Box{10, 30, 1, WHITE},
        Box{20, 20, 20, YELLOW},
    }

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The biggest one is", boxes.BiggestsColor().String())

    fmt.Println("Let's paint them all black")
    boxes.PaintItBlack()
    fmt.Println("The color of the second one is", boxes[1].color.String())

    fmt.Println("Obviously, now, the biggest one is", boxes.BiggestsColor().String()) // BiggestsColor绑定到BoxList，String绑定到Color
}


