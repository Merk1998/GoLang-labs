package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileIn, err := os.Open("stdin.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := fileIn.Close(); err != nil {
			panic(err)
		}
	}()

	stat, err := fileIn.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, stat.Size())
	_, err = fileIn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	strBuf := string(buf)
	strBufSlice := strings.Split(strBuf, " ")
	fmt.Println(strBufSlice)

	for i := 0; i < len(strBufSlice); i++ {
		if strBufSlice[i] == "" {
			strBufSlice = append(strBufSlice[:i],strBufSlice[i+1:]...)
			i--
		}
	}

	fmt.Println(strBufSlice)
	var intBufSlice []int
	for _, s := range strBufSlice {
		num, err := strconv.Atoi(s)
		if err == nil {
			intBufSlice = append(intBufSlice, num)
		} else {
			fmt.Println(err)
			return
		}
	}

	fmt.Println(intBufSlice)
	InsertionSort(intBufSlice)
	fmt.Println(intBufSlice)
}

func InsertionSort (intSlice []int) {
	for i := 1; i < len(intSlice); i++ {
		key := intSlice[i]
		lst := i
		for j := i - 1; j > -1; j-- {
			if intSlice[j] < key {
				break
			}
			intSlice[j + 1] = intSlice[j]
			lst = j
		}
		intSlice[lst] = key
	}
}