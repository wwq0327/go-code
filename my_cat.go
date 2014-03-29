package main

import "flag"
import "fmt"
import "os"
import "bufio"

/* 
 * 读取文件内容，并打印出来
 */

var infile *string = flag.String("i", "infile", "被显示的文件名")

//---------读取文件内容----------------------------------

func readFile(infile string)(lines []string, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return 
	}
	defer file.Close()

	br := bufio.NewReader(file)
	for {
		line, ok := br.ReadString('\n')
		
		if ok != nil {
			break
		} else {
			err = ok
			return 
		}
		lines = append(lines, line)
	}
	return

}

//---------------------------程序主函数-------------------------------
func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile)
	}

	lines, err := readFile(*infile)
	if err == nil {
		fmt.Println("Read valudes:", lines)
	} else {
		fmt.Println(err)
	}

}
