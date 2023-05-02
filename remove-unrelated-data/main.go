package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	delete(dataMap, key)
	return dataMap
}

func main() {
	data := map[string]any{"name": "Edo", "age": 20, "address": "Jakarta"}
	newData := removeUnrelated(data, "address")
	fmt.Println(newData)

	data2 := map[string]any{"run": "lari", "jump": "loncat", "swim": "berenang"}
	newData2 := removeUnrelated(data2, "jump")
	fmt.Println(newData2)
}
