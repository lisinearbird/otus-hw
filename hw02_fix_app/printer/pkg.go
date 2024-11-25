package printer

import (
	"fmt" //nolint:gci
	"hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		fmt.Println(staff[i].String())
	}
}
