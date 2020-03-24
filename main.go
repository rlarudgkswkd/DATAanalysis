package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var fileSave string = "10000_20_"
	num := 0

	file1, _ := os.Open("10000 20 9021.txt") // hello1.txt 파일 열기
	defer file1.Close()                      // main 함수가 끝나기 직전에 파일을 닫음

	i := 0
	filenum := 1
	//s := make([]int, 0)

	for {
		file2, _ := os.Create("C:\\Users\\khkim6\\go\\src\\github.com\\Beautykkh\\learngo\\readmyfile\\analysis\\10000-20-9021\\" + fileSave + strconv.Itoa(filenum) + ".txt")

		for j := 0; j < 20; j++ {
			fmt.Fscan(file1, &num)
			fmt.Print("i:", i)
			fmt.Print(" num:", num, "\n")
			fmt.Fprintln(file2, num)

			i++
		}

		if i%20 == 0 {
			//s = make([]int, 0)
			filenum++
			file2.Close()
		}

		if i == 1200 {
			break
		}

		/*fmt.Fscan(file1, &num)
		s = append(s, num) // 파일을 읽은 뒤 공백, 개행 문자로
		// 구분된 문자열에서 입력을 받음
		fmt.Println("i:", i)
		fmt.Println(s) // 1 1.1 Hello
		//fmt.Fprintln(file2, num)
		i++*/

	}

}
