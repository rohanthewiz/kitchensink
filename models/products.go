package models

type Product struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

var Products = []Product{
	{Label: "John Brown", Value: "john_brown"},
	{Label: "Mary Jane", Value: "mary_jane"},
	{Label: "JoAnn Smith", Value: "joann"},
	{Label: "Mike Jackson", Value: "mickey"},
	{Label: "Little John", Value: "little_john"},
}
