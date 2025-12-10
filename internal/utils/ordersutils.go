package utils

import "fmt"

func centsToEuro(amountCents int) float64 {
	return float64(amountCents) / 100
}

func PrintTotalRevenues(totalRevenues int) {
	fmt.Printf("Total revenue: %.2f EUR\n", centsToEuro(totalRevenues))
}

func PrintMarketplaceRevenues(marketplaceRevenue []MarketplaceRevenue) {
	fmt.Println("Revenues by marketplace:")

	for _, mr := range marketplaceRevenue {
		text := fmt.Sprintf("- %s: %.2f EUR", mr.MarketPlace, centsToEuro(mr.AmountCents))
		fmt.Println(text)
	}
}

func PrintSuspiciousOrders(suspiciousOrders SuspiciousOrdersMap) {
	fmt.Println("Suspicious orders:")
	for orderID, reason := range suspiciousOrders {
		text := fmt.Sprintf("- %s: %s", orderID, reason)
		fmt.Println(text)
	}
}
