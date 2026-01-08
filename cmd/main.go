package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	UserId      int      `json:"user_id"`
	Name        string   `json:"name"`
	Age         int      `json:"age"`
	Password    string   `json:"password"`
	Permissions []string `json:"roles"`
}

type personel struct {
	Users []user `json:"users"`
}

func main() {
	user1 := user{
		UserId:      1,
		Name:        "behrang",
		Age:         23,
		Password:    "jdaiuh342",
		Permissions: []string{"programmer", "admin"},
	}
	user2 := user{
		UserId:      2,
		Name:        "shahrzaad",
		Age:         28,
		Password:    "ifsudjhgi342",
		Permissions: []string{"contractor"},
	}

	personelSlice := personel{Users: []user{user1, user2}}

	f, err := os.Create("personel.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		panic(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ") // Two spaces for indentation
	err = encoder.Encode(personelSlice)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		panic(err)
	}

	fmt.Println("Successfully wrote pretty JSON to personel.json")
}
