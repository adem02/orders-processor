package orders

import (
	"fmt"
	"sort"

	"github.com/adem02/orders-processor/internal/utils"
)

type ProcessResult struct {
	TotalRevenue        int
	MarketplacesRevenue []utils.MarketplaceRevenue
	SuspiciousOrders    utils.SuspiciousOrdersMap
}

func ProcessOrdersData(orders []Order) *ProcessResult {
	if len(orders) == 0 {
		return nil
	}

	totalRevenuesCents := 0
	suspiciousOrders := make(utils.SuspiciousOrdersMap)
	revenuesIndexedByMarketplace := make(map[string]int)

	for _, order := range orders {
		if order.IsSuspicious() {
			if order.Marketplace == "" && order.AmountCents < 0 {
				suspiciousOrders[order.ID] = "empty marketplace and negative amount"
			} else if order.Marketplace == "" {
				suspiciousOrders[order.ID] = "empty marketplace"
			} else {
				suspiciousOrders[order.ID] = fmt.Sprintf("negative amount (%d)", order.AmountCents)
			}
		} else {
			totalRevenuesCents = totalRevenuesCents + order.AmountCents
			revenuesIndexedByMarketplace[order.Marketplace] += order.AmountCents
		}
	}

	marketplacesRevenue := make([]utils.MarketplaceRevenue, 0, len(revenuesIndexedByMarketplace))
	for marketplace, amountCents := range revenuesIndexedByMarketplace {
		marketplacesRevenue = append(marketplacesRevenue, utils.MarketplaceRevenue{
			MarketPlace: marketplace,
			AmountCents: amountCents,
		})
	}

	sort.Slice(marketplacesRevenue, func(i, j int) bool {
		return marketplacesRevenue[i].AmountCents > marketplacesRevenue[j].AmountCents
	})

	return &ProcessResult{
		TotalRevenue:        totalRevenuesCents,
		MarketplacesRevenue: marketplacesRevenue,
		SuspiciousOrders:    suspiciousOrders,
	}
}
