package main

import "fmt"

func howManyElements(data []any) int {
	return len(data)
}

func main() {
	data1 := []any{1, 2, 3, 4, 5}
	count1 := howManyElements(data1)
	fmt.Println(count1)

	data2 := []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	count2 := howManyElements(data2)
	fmt.Println(count2)

	data3 := []any{1, 1, 1, 5, 5, 5}
	count3 := howManyElements(data3)
	fmt.Println(count3)

	data4 := []any{"Edo", "Budi", "Joko", "Tono"}
	count4 := howManyElements(data4)
	fmt.Println(count4)

	data5 := []any{"Edo", "Budi", "Joko", "Tono", "Edo", "Budi", "Joko", "Tono"}
	count5 := howManyElements(data5)
	fmt.Println(count5)

	data6 := []any{true, false, true, false, true, false}
	count6 := howManyElements(data6)
	fmt.Println(count6)
}
