package models

type Product struct {
	Label string
	Value string
}

var Products = []Product{
	{Label: "John Brown", Value: "john_brown"},
	{Label: "Mary Jane", Value: "mary_jane"},
	{Label: "JoAnn Smith", Value: "joann"},
	{Label: "Mike Jackson", Value: "mickey"},
	{Label: "Little John", Value: "little_john"},
}
