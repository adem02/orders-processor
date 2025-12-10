package orders

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func GetOrdersFileContent(filename string) ([]Order, error) {
	data, err := os.ReadFile(filepath.Join("data", filename))
	if err != nil {
		return nil, fmt.Errorf("failed reading file %s: %w", filename, err)
	}

	ordersInput := make([]OrderInput, 0)
	if err := json.Unmarshal(data, &ordersInput); err != nil {
		return nil, fmt.Errorf("failed parsing json content of %s: %w", filename, err)
	}

	ordersList, err := ToOrdersList(ordersInput)
	if err != nil {
		return nil, fmt.Errorf("failed converting orders input to orders list %s: %w", filename, err)
	}

	return ordersList, nil
}
