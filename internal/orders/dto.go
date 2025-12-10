package orders

type OrderInput struct {
	ID          string `json:"id"`
	Marketplace string `json:"marketplace"`
	Country     string `json:"country"`
	AmountCents int    `json:"amount_cents"`
	CreatedAt   string `json:"created_at"`
}
