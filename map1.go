package main
import "fmt"

// PersonInfo是一个包含个人详细信息的类型
type PersonInfo struct {
	ID string
	Name string
	Address string
}

func main() {
	var personDB map[string] PersonInfo // myMap是声明的map变量名，string是键的类型，PersonInfo则是其中所存放的值类型。
	personDB = make(map[string] PersonInfo)

	// 往这个map里插入几条数据
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203, ..."}
	personDB["1"] = PersonInfo{"1", "jack", "Room 101, ..."}
	
	person, ok := personDB["12345"]

	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234", ", address is ", person.Address)
	} else {
		fmt.Println("Did not find person with ID 1234")
	}
}
