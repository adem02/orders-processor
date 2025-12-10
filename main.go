package main

import (
	"fmt"
	"os"

	"github.com/adem02/orders-processor/internal/orders"
	"github.com/adem02/orders-processor/internal/utils"
)

func main() {
	filename := os.Args[1]

	ordersList, err := orders.GetOrdersFileContent(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	processResult := orders.ProcessOrdersData(ordersList)

	utils.PrintTotalRevenues(processResult.TotalRevenue)
	fmt.Println()
	utils.PrintMarketplaceRevenues(processResult.MarketplacesRevenue)
	fmt.Println()
	utils.PrintSuspiciousOrders(processResult.SuspiciousOrders)
}
