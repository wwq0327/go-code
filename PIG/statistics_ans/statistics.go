package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"math"
)

// 定义HTML内容
const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Statistics</h3>
<p>Computes basic statistics for a given list of numbers</p>`
	form = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

type statistics struct {
	numbers []float64 // 切片
	mean    float64   // 平均数
	median  float64   //中位数
	mode    []float64 // 众数
	stdDev  float64 //标准差 
}

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("Failed to start server", err)
	}
}

// 主页：homePage
func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
			fmt.Fprint(writer, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}

	fmt.Fprintf(writer, pageBottom)
}

// 数据处理模块
func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1) // 将用户输入的逗号转换为空格， －1，表示全部都进行处理
		for _, field := range strings.Fields(text) {    // 使用Fields将内容以空格为界切将字符串切成切片
			if x, err := strconv.ParseFloat(field, 64); err != nil { // 转换成float64类型
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}

	if len(numbers) == 0 {
		return numbers, "", false
	}

	return numbers, "", true
}

// 输出统计结果
func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
<tr><td>Mode</td><td>%f</td></tr>
<tr><td>Std.Dev</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median, stats.mode, stats.stdDev)
}

// 接收一个numbers 有切片，反回一个statistics 结构
func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)                      // sort.Float64s 对numbers进行一个原地升序的排序工作
	stats.mean = sum(numbers) / float64(len(numbers)) // 求平均数
	stats.median = median(numbers)     //求中位数
	stats.mode = mode(numbers)
	stats.stdDev = stdDev(numbers, stats.mean)
	return stats
}

// 求和，并返回总和
func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

// 求中位数
func median(numbers []float64) float64 {
	middle := len(numbers) / 2 // 直接取整，小数位会被自动去掉
	result := numbers[middle]
	if len(numbers)%2 == 0 { //如果数据个数为偶数个时，取中间两数的平均值
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

// 众数
func mode(numbers []float64) (modes []float64) {
	frequencies := make(map[float64]int, len(numbers))
	highestFrequency := 0
	for _, x := range numbers {
		frequencies[x]++
		if frequencies[x] > highestFrequency {
			highestFrequency = frequencies[x]
		}
	}
	for x, frequency := range frequencies {
		if frequency == highestFrequency {
			modes = append(modes, x)
		}
	}
	if highestFrequency == 1 || len(modes) == len(frequencies) {
		modes = modes[:0]
	}
	sort.Float64s(modes)
	return modes
}

// 标准差
func stdDev(numbers []float64, mean float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += math.Pow(number-mean, 2)
	}
	variance := total / float64(len(numbers)-1)
	return math.Sqrt(variance)
}
