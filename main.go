package main

import (
	"fmt"

	"github.com/csmyx/LearnGo/LinkTable"
)

func main() {
	nums := []int{2, 1, 2, 2, 5, 3, 5, 7}
	lt := LinkTable.NewLinkTable()
	for _, num := range nums {
		lt.InsertVal(num)
	}
	fmt.Println(lt.ToString())
}
