package export

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadFile(path string) (*Model, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}

	var model Model
	if err := json.NewDecoder(f).Decode(&model); err != nil {
		return nil, fmt.Errorf("reading json: %w", err)
	}

	return &model, nil
}
