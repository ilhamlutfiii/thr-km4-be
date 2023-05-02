package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		return append([]int{newData}, data...)
	} else if position == "down" {
		return append(data, newData)
	} else {
		panic("Posisi tidak diketahui, pilih antara \"up\" atau \"down\"")
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	newData := 6

	newArrayData := AddElement(data, newData, "up")
	fmt.Println(newArrayData)

	newArrayData = AddElement(data, newData, "down")
	fmt.Println(newArrayData)
}
