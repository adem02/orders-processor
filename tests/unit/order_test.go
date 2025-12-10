package orders

import (
	"testing"
	"time"

	"github.com/adem02/orders-processor/internal/orders"
)

func TestNewOrderWillSucceed(t *testing.T) {
	id := "order-1"
	marketplace := "amazon"
	country := "FR"
	amountCents := 12345
	createdAt := "2024-11-01T10:15:00Z"

	o, err := orders.NewOrder(id, marketplace, country, amountCents, createdAt)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if o.ID != id {
		t.Errorf("expected ID %q, got %q", id, o.ID)
	}
	if o.Marketplace != marketplace {
		t.Errorf("expected Marketplace %q, got %q", marketplace, o.Marketplace)
	}
	if o.Country != country {
		t.Errorf("expected Country %q, got %q", country, o.Country)
	}
	if o.AmountCents != amountCents {
		t.Errorf("expected AmountCents %d, got %d", amountCents, o.AmountCents)
	}

	expectedTime, _ := time.Parse(time.RFC3339, createdAt)
	if !o.CreatedAt.Equal(expectedTime) {
		t.Errorf("expected CreatedAt %v, got %v", expectedTime, o.CreatedAt)
	}
}

func TestNewOrderInvalidDateWillFail(t *testing.T) {
	_, err := orders.NewOrder("id", "mp", "FR", 100, "not-a-date")
	if err == nil {
		t.Fatalf("expected error for invalid date")
	}
}

func TestOrderIsSuspicious(t *testing.T) {
	tests := []struct {
		name  string
		order orders.Order
	}{
		{
			name: "negative amount",
			order: orders.Order{
				ID:          "1",
				Marketplace: "amazon",
				AmountCents: -100,
			},
		},
		{
			name: "empty marketplace",
			order: orders.Order{
				ID:          "2",
				Marketplace: "",
				AmountCents: 100,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.order.IsSuspicious() {
				t.Fatalf("expected order to be suspicious")
			}
		})
	}
}

func TestOrderIsNotSuspicious(t *testing.T) {
	tests := []struct {
		name  string
		order orders.Order
	}{
		{
			name: "zero amount with valid marketplace",
			order: orders.Order{
				ID:          "3",
				Marketplace: "amazon",
				AmountCents: 0,
			},
		},
		{
			name: "positive amount and marketplace",
			order: orders.Order{
				ID:          "4",
				Marketplace: "amazon",
				AmountCents: 100,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.order.IsSuspicious() {
				t.Fatalf("expected order to NOT be suspicious")
			}
		})
	}
}

func TestToOrdersListWillSucceed(t *testing.T) {
	inputs := []orders.OrderInput{
		{
			ID:          "1",
			Marketplace: "amazon",
			Country:     "FR",
			AmountCents: 100,
			CreatedAt:   "2024-11-01T10:15:00Z",
		},
		{
			ID:          "2",
			Marketplace: "ebay",
			Country:     "DE",
			AmountCents: 200,
			CreatedAt:   "2024-11-02T10:15:00Z",
		},
	}

	orders, err := orders.ToOrdersList(inputs)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(orders) != len(inputs) {
		t.Fatalf("expected %d orders, got %d", len(inputs), len(orders))
	}

	for i, in := range inputs {
		if orders[i].ID != in.ID {
			t.Errorf("order[%d].ID = %q, expected %q", i, orders[i].ID, in.ID)
		}
		if orders[i].Marketplace != in.Marketplace {
			t.Errorf("order[%d].Marketplace = %q, expected %q", i, orders[i].Marketplace, in.Marketplace)
		}
		if orders[i].Country != in.Country {
			t.Errorf("order[%d].Country = %q, expected %q", i, orders[i].Country, in.Country)
		}
		if orders[i].AmountCents != in.AmountCents {
			t.Errorf("order[%d].AmountCents = %d, expected %d", i, orders[i].AmountCents, in.AmountCents)
		}
	}
}

func TestToOrdersListWillFail(t *testing.T) {
	inputs := []orders.OrderInput{
		{
			ID:          "1",
			Marketplace: "amazon",
			Country:     "FR",
			AmountCents: 100,
			CreatedAt:   "2024-11-01T10:15:00Z",
		},
		{
			ID:          "2",
			Marketplace: "ebay",
			Country:     "DE",
			AmountCents: 200,
			CreatedAt:   "invalid-date",
		},
	}

	result, err := orders.ToOrdersList(inputs)
	if err == nil {
		t.Fatalf("expected error for invalid OrderInput")
	}

	if result != nil {
		t.Fatalf("expected no result on error, got %#v", result)
	}
}
