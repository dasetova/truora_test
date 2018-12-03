package models

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Recipe struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Indications string   `json:"indications"`
	CategoryID  int      `json:"category_id"`
	Category    Category `json:"category"`

	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
	MeasureUnit string  `json:"measure_unit"`
	RecipeID    int     `json:"recipe_id"`
	//Recipe      Recipe  `json:"recipe"`
}
