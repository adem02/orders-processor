package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/adem02/orders-processor/internal/orders"
	"github.com/adem02/orders-processor/internal/utils"
)

func main() {
	filter := flag.String("from", "", "Filter date in format YYYY-MM-DD")
	flag.Parse()

	agrs := flag.Args()
	if len(agrs) < 1 {
		log.Fatal("Data file argument is required")
	}
	filename := agrs[0]
	var from *time.Time

	if *filter != "" {
		parsedTime, err := time.Parse("2006-01-02", *filter)
		if err != nil {
			log.Fatal("Invalid date format. Please use YYYY-MM-DD.")
		}

		from = &parsedTime
	}

	ordersList, err := orders.GetOrdersFileContent(filename, from)
	if err != nil {
		log.Fatal(err.Error())
	}

	processResult := orders.ProcessOrdersData(ordersList)

	if processResult == nil {
		fmt.Println("No orders to process")
	} else {
		utils.PrintTotalRevenues(processResult.TotalRevenue)
		fmt.Println()
		utils.PrintMarketplaceRevenues(processResult.MarketplacesRevenue)
		fmt.Println()
		utils.PrintSuspiciousOrders(processResult.SuspiciousOrders)
	}
}
