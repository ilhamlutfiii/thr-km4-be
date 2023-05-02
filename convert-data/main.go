package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	userFields := strings.Split(data, ",")
	age, _ := strconv.Atoi(userFields[1])
	user := User{
		Name:    userFields[0],
		Age:     age,
		Address: userFields[2],
	}

	return user
}

func main() {
	user1 := ConvertData("Edo,20,Jakarta")
	fmt.Printf("name: %s, age: %d, address: %s\n", user1.Name, user1.Age, user1.Address)

	user2 := ConvertData("Budi,30,Bandung")
	fmt.Printf("name: %s, age: %d, address: %s\n", user2.Name, user2.Age, user2.Address)

}
