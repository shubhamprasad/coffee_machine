package usecases

import (
	"coffee_machine/models"
	"fmt"
	"sync"
)

// readInput - function to parse the input json mapped as a struct
func readInput(inputStruct models.InputStruct) (outlets int, ingredientsList map[string]models.Ingredient, beverageList []models.Beverage) {

	outlets = int(inputStruct.Machine[Outlets].(map[string]interface{})[CountN].(float64))
	ingredientsList = make(map[string]models.Ingredient)

	for ingredient, value := range inputStruct.Machine[TotalItemsQuantity].(map[string]interface{}) {
		ingredientsList[ingredient] = models.Ingredient{ingredient, int(value.(float64))}
	}

	ingredient := models.Ingredient{}
	beverage := models.Beverage{}
	for beverageName, value := range inputStruct.Machine[Beverages].(map[string]interface{}) {
		beverage.Name = beverageName
		for ingredientName, ingredientValue := range value.(map[string]interface{}) {
			ingredient.Name = ingredientName
			ingredient.Quantity = int(ingredientValue.(float64))
			beverage.IngredientsList = append(beverage.IngredientsList, ingredient)
			ingredient = models.Ingredient{}
		}
		beverageList = append(beverageList, beverage)
		beverage = models.Beverage{}
	}

	return
}

// sendBeverageOrders - sends the beverages requested to a channel
func sendBeverageOrders(beverageChannel chan<- models.Beverage, beverageList []models.Beverage) {
	for _, beverage := range beverageList {
		beverageChannel <- beverage
	}
	close(beverageChannel)
}

// prepareBeverage - prepares beverages after checking the ingredients
func prepareBeverage(coffeeMachine *models.CoffeeMachine, beverageChannel <-chan models.Beverage, wg *sync.WaitGroup) {

	defer wg.Done()

	for beverage := range beverageChannel {
		checkAndConsumeIngredients(coffeeMachine, beverage)
	}
}

// checkAndConsumeIngredients - checks if ingredients are sufficient and then prepares them
func checkAndConsumeIngredients(coffeeMachine *models.CoffeeMachine, beverage models.Beverage) {

	message := ""
	areIngredientsLess := false

	// Below is the critical section. Mutex is used to avoid race conditions
	mutex.Lock()

	// First check if all the ingredients are present. If any of them is insufficient then discard the order
	for _, beverageIngredient := range beverage.IngredientsList {
		if coffeeMachineIngredient, ok := coffeeMachine.Ingredients[beverageIngredient.Name]; ok {
			if coffeeMachineIngredient.Quantity < beverageIngredient.Quantity {
				message = fmt.Sprintf("%s cannot be prepared because %s is not sufficient", beverage.Name, beverageIngredient.Name)
				areIngredientsLess = true
				break
			}
		} else {
			message = fmt.Sprintf("%s cannot be prepared because %s is not available", beverage.Name, beverageIngredient.Name)
			areIngredientsLess = true
			break
		}
	}

	// I all ingredients are present then consume ingredients and decrease the quantity from coffee machine
	if !areIngredientsLess {
		for _, beverageIngredient := range beverage.IngredientsList {
			if coffeeMachineIngredient, ok := coffeeMachine.Ingredients[beverageIngredient.Name]; ok {
				if coffeeMachineIngredient.Quantity >= beverageIngredient.Quantity {
					coffeeMachine.Ingredients[beverageIngredient.Name] = updateIngredients(coffeeMachine.Ingredients[beverageIngredient.Name], beverageIngredient.Quantity)
				}
			}
		}
	}

	// Critical section ends
	mutex.Unlock()

	if !areIngredientsLess {
		fmt.Println(beverage.Name + " is prepared")
	} else {
		fmt.Println(message)
	}
}

// updateIngredients - updates the ingredients' quantity in the coffee machine
func updateIngredients(ingredient models.Ingredient, quantity int) models.Ingredient {
	ingredient.Quantity = ingredient.Quantity - quantity
	return ingredient
}