package reader

import (
	"encoding/json"
	"fmt"
	"hw02_fix_app/types" //nolint:gci
	"io"
	"os" //nolint:gci
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)

	return data, err
}
