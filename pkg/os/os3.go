package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	// 改变文件模式
	chmod()
	// 改变工作目录
	// chdir()
	// Chown()
	chown()
}

func chdir() {
	// Chdir
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The current directory is: %s\n", pwd)

	if err = os.Chdir("/Users/wyatt/golang/go-code/pkg"); err != nil {
		fmt.Println(err)
		return
	}

	pwd, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Now the current directory is: %s\n", pwd)
}

func chmod() {
	// test.txt's mode is: -rw-r--r--(644)
	// test.txt's mode is: -rwxrwxrwx(777)
	fi, err := os.Stat("test.txt")
	if err != nil {
		fmt.Printf("Error :%s\n", err)
		return
	}

	fmt.Printf("test.txt's mode is: %s(%o)\n", fi.Mode(), fi.Mode()&0777)

	err = os.Chmod("test.txt", 0777)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fi, err = os.Stat("test.txt")
	if err != nil {
		fmt.Printf("Error :%s\n", err)
		return
	}

	fmt.Printf("test.txt's mode is: %s(%o)\n", fi.Mode(), fi.Mode()&0777)

}

// 修改一个文件的所属用户和组
func chown() {
	fi, err := os.Stat("test.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("test.txt: uid=%d, gid=%d\n", fi.Sys().(*syscall.Stat_t).Uid, fi.Sys().(*syscall.Stat_t).Gid)

	fi, err = os.Chown("test.txt", 99, 99)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fi, err = os.Stat("test.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("test.txt: uid=%d, gid=%d\n", fi.Sys().(*syscall.Stat_t).Uid, fi.Sys().(*syscall.Stat_t).Gid)

}
