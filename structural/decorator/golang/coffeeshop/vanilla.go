package coffeeshop

// Vanilla decorator struct to add Vanilla to coffee
type Vanilla struct {
	Coffee Coffee
}

// GetCost get the cost of base coffee and plus the cost of Vanilla complement
func (m *Vanilla) GetCost() float32 {
	coffeePrice := m.Coffee.GetCost()
	return coffeePrice + 0.75
}

// GetIngredients get the ingrtedients and add the Vanilla ingredient
func (m *Vanilla) GetIngredients() string {
	coffeeIngredients := m.Coffee.GetIngredients()
	return coffeeIngredients + ", Vanilla"
}

// GetTax get the tax with Vanilla
func (m *Vanilla) GetTax() float32 {
	return 0.1 * m.GetCost()
}
