package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	UserId      int      `json: "user_id"`
	Name        string   `json: "name"`
	Age         int      `json: "age"`
	Password    string   `json: "password"`
	Permissions []string `json: "roles"`
}

func main() {

	// b, err := os.ReadFile("jsonOutput.json")
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	f, err := os.Open("jsonOutput.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer f.Close()

	var person user

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&person)
	fmt.Println(person)

}
