package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var fileSave string = "9000_20_"
	num := 0.0
	arr1 := make([]float64, 0)
	var averRage float64

	//set the variables
	diviNum := 20.0 //몇으로 나눌지
	setPartialNum := 20
	setWholeNum := 1200
	theOpenWay := "C:\\Users\\khkim6\\Desktop\\testingcase\\9000_20_7255.txt"
	theCreateWay := "C:\\Users\\khkim6\\Desktop\\testingcase\\90000_20_7225\\"
	theSaveAvgWay := "C:\\Users\\khkim6\\Desktop\\testingcase\\aveRage9000_20.txt"

	file1, _ := os.Open(theOpenWay) // hello1.txt 파일 열기
	defer file1.Close()             // main 함수가 끝나기 직전에 파일을 닫음

	i := 0
	filenum := 1
	//s := make([]int, 0)

	for {
		file2, _ := os.Create(theCreateWay + fileSave + strconv.Itoa(filenum) + ".txt")

		for j := 0; j < setPartialNum; j++ {
			fmt.Fscan(file1, &num)
			fmt.Print("i:", i)
			fmt.Print(" num:", num, "\n")
			fmt.Fprintln(file2, num)
			averRage = averRage + num

			i++

		}

		if i%setPartialNum == 0 {
			//s = make([]int, 0)
			arr1 = append(arr1, averRage/diviNum) //calculate the average
			averRage = 0.0
			filenum++
			file2.Close()
		}

		if i == setWholeNum {
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
	var selection string

	fmt.Printf("Do you want to save? -> Y OR N :")
	fmt.Scanf("%s", &selection)

	if selection == "Y" {
		file3, _ := os.Create(theSaveAvgWay)
		defer file3.Close()
		for _, num := range arr1 {
			fmt.Fprintln(file3, num)
		}
	}

	fmt.Println(arr1)

}
