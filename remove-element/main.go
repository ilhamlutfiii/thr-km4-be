package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	if position == "left" {
		return arr[1:]
	} else if position == "right" {
		return arr[:len(arr)-1]
	}
	return arr
}

func main() {
	arr1 := []any{1, 2, 3, 4, 5}
	result1 := removeLeftRight(arr1, "left")
	fmt.Println(result1)

	arr2 := []any{1, 2, 3, 4, 5}
	result2 := removeLeftRight(arr2, "right")
	fmt.Println(result2)

	arr3 := []any{"Edo", "Budi", "Joko", "Tono"}
	result3 := removeLeftRight(arr3, "left")
	fmt.Println(result3)

	arr4 := []any{"Edo", "Budi", "Joko", "Tono"}
	result4 := removeLeftRight(arr4, "right")
	fmt.Println(result4)
}
