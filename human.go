package main
import "fmt"
import "strconv"

type Human struct {
	name string
	age int
}

type Student struct {
	Human // 匿名字段
	number int
}

func (s Student) String() string {
	return "<" + s.name + "-" + strconv.Itoa(s.age) + "-" + strconv.Itoa(s.number) + ">"
}

func main() {
	s := Student{Human{"wyatt", 33}, 97019}
	fmt.Println("this name: ", s.name)
	fmt.Println("this human name:", s.Human.name)
	fmt.Println("ALL INFO", s)
}
