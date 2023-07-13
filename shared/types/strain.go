package types

type Strain struct {
	ID                string   `json:"id"`
	ProducerID        string   `json:"producer_id"`
	Name              string   `json:"name"`
	Category          string   `json:"category"`
	Genetics          string   `json:"genetics"`
	THC               float64  `json:"thc"`
	Terpenes          float64  `json:"terpenes"`
	CBD               float64  `json:"cbd"`
	TotalCannabinoids float64  `json:"total_cannabinoids"`
	TerpeneList       []string `json:"terpene_list"`
	Price             float64  `json:"price"`
	HarvestDate       string   `json:"harvest_date"`
	Images            []string `json:"images"`
}
