package main
import "fmt"

type Studenter interface {
	GetName()
}

type Student struct {
	name string
}

func (s Student) GetName() {
	fmt.Println("Name:", s.name)
}

func main() {
	var a Studenter
	s := Student{"wyatt"}
	a = s
	a.GetName()
}
