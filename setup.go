package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func makeRange(from, to int) []int {
	arr := make([]int, to-from+1)
	for i := range arr {
		arr[i] = from + i
	}

	return arr
}

func main() {
	days := makeRange(1, 24)
	var builder strings.Builder

	for _, day := range days {

		builder.WriteString("day")
		builder.WriteString(strconv.Itoa(day))

		directory := filepath.Join(".", builder.String())
		err := os.MkdirAll(directory, os.ModePerm)

		if err == nil {
			_, err := os.Create(filepath.Join(directory, "/phase-1.go"))
			if err != nil {
				fmt.Println(err)
			}

			_, err = os.Create(filepath.Join(directory, "/phase-2.go"))
			if err != nil {
				fmt.Println(err)
			}
		}
		builder.Reset()
	}
}
