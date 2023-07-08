package types

type Catalog struct {
	Producers []Producer `yaml:"producers"`
}

type Producer struct {
	ID       string    `yaml:"id"`
	Name     string    `yaml:"name"`
	Location string    `yaml:"location"`
	Website  string    `yaml:"website"`
	Phone    string    `yaml:"phone"`
	Email    string    `yaml:"email"`
	Products []Product `yaml:"products"`
}

type Product struct {
	ID    string  `yaml:"id"`
	Name  string  `yaml:"name"`
	Type  string  `yaml:"type"`
	THC   float64 `yaml:"thc"`
	Terps float64 `yaml:"terps"`
	Price float64 `yaml:"price"`
}
