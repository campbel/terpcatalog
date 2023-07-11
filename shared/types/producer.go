package types

type Producer struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Website string          `json:"website"`
	Email   string          `json:"email"`
	Phone   string          `json:"phone"`
	Address string          `json:"address"`
	Contact ProducerContact `json:"contact"`
}

type ProducerContact struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
