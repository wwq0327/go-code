package main
import "fmt"

/*
 *一个博客结构
 */

// Blog是一个含有blog信息的类型
type Blog struct {
	ID string             // blog.id
	Title string          // blog title
	Author string         // author
	Content string        // blog content
	Datetime string       // created datetime
}

// .......................................................

func main() {
	var blogs map[string] Blog  // blogs 为一个map变量名，键为string，值为Blog
	blogs = make(map[string] Blog)

	// add data
	blogs["1"] = Blog{"1", "hello, world", "wyatt", "Mr. Hello", "2014-03-28"}
	blogs["2"] = Blog{"2", "Go blog", "wyatt", "this is a blog by Golang", "2014-03-28"}
	
	// print blog info
	for _, blog := range blogs {
		fmt.Printf("# %s %s\n", blog.ID, blog.Title)
		fmt.Println("========================")
		fmt.Printf("Author: %s Publish Date: %s\nContent: %s\n", blog.Author, blog.Datetime, blog.Content)
		fmt.Println()

	}
}
