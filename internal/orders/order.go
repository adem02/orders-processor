package orders

import "time"

type Order struct {
	ID          string
	Marketplace string
	Country     string
	AmountCents int
	CreatedAt   time.Time
}

func NewOrder(id, marketplace, country string, amountCents int, createdAt string) (*Order, error) {
	createdAtDate, err := time.Parse(time.RFC3339, createdAt)

	if err != nil {
		return nil, err
	}

	return &Order{
		ID:          id,
		Marketplace: marketplace,
		Country:     country,
		AmountCents: amountCents,
		CreatedAt:   createdAtDate,
	}, nil
}

func (oi OrderInput) ToOrder() (*Order, error) {
	return NewOrder(oi.ID, oi.Marketplace, oi.Country, oi.AmountCents, oi.CreatedAt)
}

func (o Order) IsSuspicious() bool {
	return o.AmountCents < 0 || o.Marketplace == ""
}

func ToOrdersList(oiList []OrderInput) ([]Order, error) {
	orders := make([]Order, 0, len(oiList))

	for _, oi := range oiList {
		order, err := oi.ToOrder()

		if err != nil {
			return nil, err
		}
		orders = append(orders, *order)
	}

	return orders, nil
}

func ToFilteredOrdersList(oi []OrderInput, from time.Time) ([]Order, error) {
	orders := make([]Order, 0, len(oi))

	for _, orderInput := range oi {
		order, err := orderInput.ToOrder()
		if err != nil {
			return nil, err
		}

		if !order.CreatedAt.Before(from) {
			orders = append(orders, *order)
		}
	}

	return orders, nil
}

func CreateOrdersList(oiList []OrderInput, from *time.Time) ([]Order, error) {
	if from == nil {
		return ToOrdersList(oiList)
	}

	return ToFilteredOrdersList(oiList, *from)
}
