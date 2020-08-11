package models


/* Beverage - Struct to represent a Beverage
   Has Name and a list of Ingredients required to prepare it
*/
type Beverage struct {
	Name string
	IngredientsList []Ingredient
}
