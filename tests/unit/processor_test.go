package orders

import (
	"testing"
	"time"

	"github.com/adem02/orders-processor/internal/orders"
	"github.com/adem02/orders-processor/internal/utils"
)

func TestProcessOrdersDataWillSucceed(t *testing.T) {
	createdAt, _ := time.Parse(time.RFC3339, "2024-11-01T10:15:00Z")

	ordersList := []orders.Order{
		{
			ID:          "1",
			Marketplace: "amazon",
			AmountCents: 1000,
			CreatedAt:   createdAt,
		},
		{
			ID:          "2",
			Marketplace: "amazon",
			AmountCents: 500,
			CreatedAt:   createdAt,
		},
		{
			ID:          "3",
			Marketplace: "ebay",
			AmountCents: 2000,
			CreatedAt:   createdAt,
		},
		{
			ID:          "4",
			Marketplace: "",
			AmountCents: 300,
			CreatedAt:   createdAt,
		},
		{
			ID:          "5",
			Marketplace: "fnac",
			AmountCents: -100,
			CreatedAt:   createdAt,
		},
	}

	result := orders.ProcessOrdersData(ordersList)

	expectedTotal := 1000 + 500 + 2000
	if result.TotalRevenue != expectedTotal {
		t.Fatalf("expected total revenue %d, got %d", expectedTotal, result.TotalRevenue)
	}

	if len(result.SuspiciousOrders) != 2 {
		t.Fatalf("expected 2 suspicious orders, got %d", len(result.SuspiciousOrders))
	}

	if _, ok := result.SuspiciousOrders["4"]; !ok {
		t.Errorf("expected order 4 to be suspicious")
	}
	if _, ok := result.SuspiciousOrders["5"]; !ok {
		t.Errorf("expected order 5 to be suspicious")
	}

	expected := []utils.MarketplaceRevenue{
		{MarketPlace: "ebay", AmountCents: 2000},
		{MarketPlace: "amazon", AmountCents: 1500},
	}

	if len(result.MarketplacesRevenue) != len(expected) {
		t.Fatalf(
			"expected %d marketplace revenues, got %d",
			len(expected),
			len(result.MarketplacesRevenue),
		)
	}

	for i, exp := range expected {
		got := result.MarketplacesRevenue[i]
		if got.MarketPlace != exp.MarketPlace || got.AmountCents != exp.AmountCents {
			t.Errorf(
				"marketplace[%d] = %+v, expected %+v",
				i,
				got,
				exp,
			)
		}
	}
}

func TestProcessOrdersDataSuspiciousOrdersReasonsWillSucceed(t *testing.T) {
	createdAt, _ := time.Parse(time.RFC3339, "2024-11-01T10:15:00Z")

	ordersList := []orders.Order{
		{
			ID:          "1",
			Marketplace: "",
			AmountCents: 100,
			CreatedAt:   createdAt,
		},
		{
			ID:          "2",
			Marketplace: "amazon",
			AmountCents: -200,
			CreatedAt:   createdAt,
		},
		{
			ID:          "3",
			Marketplace: "",
			AmountCents: -300,
			CreatedAt:   createdAt,
		},
	}

	result := orders.ProcessOrdersData(ordersList)

	if len(result.SuspiciousOrders) != 3 {
		t.Fatalf("expected 3 suspicious orders, got %d", len(result.SuspiciousOrders))
	}

	expected := "empty marketplace"
	got := result.SuspiciousOrders["1"]

	if got != expected {
		t.Errorf("order 1 reason = %q, expected %q", got, expected)
	}

	expected = "negative amount (-200)"
	got = result.SuspiciousOrders["2"]

	if got != expected {
		t.Errorf("order 2 reason = %q, expected %q", got, expected)
	}

	got = result.SuspiciousOrders["3"]
	expected = "empty marketplace and negative amount"

	if got != expected {
		t.Errorf("order 3 reason = %q, expected %q", got, expected)
	}
}
