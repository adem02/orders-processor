package orders

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func GetOrdersFileContent(filename string, from *time.Time) ([]Order, error) {
	data, err := os.ReadFile(filepath.Join("data", filename))
	if err != nil {
		return nil, fmt.Errorf("failed reading file %s: %w", filename, err)
	}

	ordersInput := make([]OrderInput, 0)
	if err := json.Unmarshal(data, &ordersInput); err != nil {
		return nil, fmt.Errorf("failed parsing json content of %s: %w", filename, err)
	}

	ordersList, err := CreateOrdersList(ordersInput, from)
	if err != nil {
		return nil, fmt.Errorf("failed converting orders input to orders list %s: %w", filename, err)
	}

	return ordersList, nil
}
