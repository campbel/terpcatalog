package types

type Order struct {
	ID          string      `json:"id"`
	Information OrderInfo   `json:"information"`
	Items       []OrderItem `json:"items"`
}

type OrderInfo struct {
	Company       string       `json:"company_name"`
	LicenseNumber string       `json:"license_number"`
	Email         string       `json:"email"`
	Phone         string       `json:"phone"`
	Address       OrderAddress `json:"address"`
}

type OrderAddress struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Postal string `json:"postal"`
}

type OrderItem struct {
	StrainID string `json:"strain_id"`
	Quantity int    `json:"quantity"`
}
