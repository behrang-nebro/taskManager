package main

import (
	"fmt"

	"github.com/behrang-nebro/taskManager/internals/iofile"
)

func main() {
	// user1 := User{
	// 	UserId: 1,
	// 	Name:   "behrang",
	// 	Age:    33,
	// }
	// user2 := User{
	// 	UserId: 2,
	// 	Name:   "shahrzad",
	// 	Age:    28,
	// }
	// personelSlice := []User{user1, user2}
	// personels := personel{Users: personelSlice}

	_, err := iofile.CreateFile("personel33.json")
	if err != nil {
		fmt.Println(err)
	}
}
