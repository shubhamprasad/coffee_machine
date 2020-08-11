package models

import "fmt"


/* CoffeeMachine - Struct to represent a CoffeeMachine
   Outlets - Number of outlets which can serve in parallel
   Ingredients - A Map to keep track of Ingredients
*/
type CoffeeMachine struct {
	Outlets int
	Ingredients map[string]Ingredient
}

// GetLowRunningIngredients - Function to get low running ingredients than a particular threshold
func (c CoffeeMachine) GetLowRunningIngredients(threshold int) {
	for _, ingredient := range c.Ingredients {
		if ingredient.Quantity <= threshold {
			fmt.Println(ingredient.Name + " is running low, quantity avaialble: ", ingredient.Quantity)
		}
	}
}

// AddIngredients - Function to add ingredients to a coffee machine
func (c *CoffeeMachine) AddIngredients(ingredientName string, quantity int) {
	if _, ingredientExists := c.Ingredients[ingredientName]; ingredientExists {
		c.Ingredients[ingredientName] = addQuantity(c.Ingredients[ingredientName], quantity)
	} else {
		c.Ingredients[ingredientName] = Ingredient{ingredientName, quantity}
	}
}

// addQuantity - private member function to add quantity
func addQuantity(ingredient Ingredient, quantity int) Ingredient {
	ingredient.Quantity = ingredient.Quantity + quantity
	return ingredient
}