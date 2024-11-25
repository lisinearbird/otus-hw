package main

import (
	"fmt" //nolint:gci
	"hw02_fix_app/printer"
	"hw02_fix_app/reader"
	"hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")

	fmt.Scanln(&path)

	if len(path) == 0 {
		path = "data.json"
	}

	var staff []types.Employee

	staff, err := reader.ReadJSON(path)
	if err != nil {
		fmt.Println("Error:", err)
	}

	printer.PrintStaff(staff)
}
