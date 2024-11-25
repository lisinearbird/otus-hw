package main

import (
	"fmt"

	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")

	_, err := fmt.Scanln(&path)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if len(path) == 0 {
		path = "data.json"
	}

	var staff []types.Employee

	staff, err = reader.ReadJSON(path, -1)
	if err != nil {
		fmt.Println("Error:", err)
	}

	printer.PrintStaff(staff)
}
