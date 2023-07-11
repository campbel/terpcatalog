package types

type Strain struct {
	ID          string   `json:"id"`
	ProducerID  string   `json:"producer_id"`
	Name        string   `json:"name"`
	Category    string   `json:"category"`
	Genetics    string   `json:"genetics"`
	THC         float64  `json:"thc"`
	Terpenes    float64  `json:"terpenes"`
	Price       float64  `json:"price"`
	HarvestDate string   `json:"harvest_date"`
	Images      []string `json:"images"`
}
