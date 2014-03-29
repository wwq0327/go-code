package main
import "fmt"

type Human struct {
	name string
	age int
}

type Student struct {
	Human
	number int
}

func main() {
	s := Student{Human{"wyatt", 33}, 97019}
	fmt.Println("this name: ", s.name)
	fmt.Println("this human name:", s.Human.name)
}
