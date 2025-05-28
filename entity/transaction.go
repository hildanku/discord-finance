package entity

type Transaction struct {
	ID        string `json:"id,omitempty"`
	UserID    string `json:"user_id"`
	Type      string `json:"type"`
	Amount    int    `json:"amount"`
	Note      string `json:"note"`
	CreatedAt string `json:"created_at"`
}
