package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Song struct {
	Title    string
	Filename string
	Seconds  int
}

type SongsSlice struct {
	Songs []Song
}

func main() {
	if len(os.Args) == 1 || !strings.HasSuffix(os.Args[1], ".m3u") {
		fmt.Printf("usage: %s <file.m3u>\n", filepath.Base(os.Args[0])) // filepath.Base(path string)返回一个只有文件名的字符串
		os.Exit(1)
	}
	if rawByte, err := ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		songs := readM3uPlaylist(string(rawByte))
		// writePisPlaylist(songs)
		writePisPlaylistToJson(songs)
	}
}

func readM3uPlaylist(data string) (songs []Song) {
	var song Song
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#EXTM3U") {
			continue
		}

		if strings.HasPrefix(line, "#EXTINF:") {
			song.Title, song.Seconds = parseExtinfLine(line)
		} else {
			song.Filename = strings.Map(mapPlatformDirSeparator, line)
		}

		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{} // 将song结构数据清零
		}
	}

	return songs
}

// 处理行信息
// 分析出标题及时长
// 返回处理之后的标题及时长
func parseExtinfLine(line string) (title string, seconds int) {
	if i := strings.IndexAny(line, "-0123456789"); i > -1 { // 查找第一个数字或负号的位置， 如果为－1，则表示未找到
		const separator = ","
		line = line[i:]
		if j := strings.Index(line, separator); j > -1 { // Index()，获取,第一出现的位置，如果为－1，则表示未找到
			title = line[j+len(separator):] // 标题内容，对字符串进行切片处理
			var err error
			if seconds, err = strconv.Atoi(line[:j]); err != nil {
				log.Printf("failed to read the duration for '%s': %v\n", title, err)
				seconds = -1
			}
		}
	}

	return title, seconds
}

func mapPlatformDirSeparator(char rune) rune {
	if char == '/' || char == '\\' {
		return filepath.Separator
	}

	return char
}

// func writePisPlaylist(songs []Song) {
// 	fmt.Println("[playlist]")
// 	for i, song := range songs {
// 		i++
// 		fmt.Printf("File%d=%s\n", i, song.Filename)
// 		fmt.Printf("Title%d=%s\n", i, song.Title)
// 		fmt.Printf("Length%d=%d\n", i, song.Seconds)
// 	}

// 	fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
// }

func writePisPlaylistToJson(songs []Song) {
	var s SongsSlice
	for _, song := range songs {
		s.Songs = append(s.Songs, Song{song.Filename, song.Title, song.Seconds})
	}

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}

	fmt.Println(string(b))
}
