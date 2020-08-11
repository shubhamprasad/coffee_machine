package usecases

import (
	"coffee_machine/models"
	"fmt"
	"sync"
)

// InitializeCoffeeMachine - Creates a coffee machine and loads ingredients in it
func InitializeCoffeeMachine(inputStruct models.InputStruct) (models.CoffeeMachine, []models.Beverage) {

	defer func() () {
		if err := recover(); err != nil {
			fmt.Println("Panic occured in InitializeCoffeeMachine")
		}

	}()

	outlets, ingredientsList, beverageList := readInput(inputStruct)

	coffeeMachine := models.CoffeeMachine{
		outlets, ingredientsList,
	}

	return coffeeMachine, beverageList
}

// PrepareBeverages - Takes orders for beverages
func PrepareBeverages(coffeeMachine *models.CoffeeMachine, beverageList []models.Beverage) {

	beverageChannel := make(chan models.Beverage)
	var wg sync.WaitGroup

	if coffeeMachine.Outlets == 0 {
		fmt.Println("No outlets in Coffee Machine")
		return
	}

	// Send orders to a channel concurrently
	go sendBeverageOrders(beverageChannel, beverageList)

	/* Initialize workers to prepare beverages concurrently. Number of workers running in parallel
	   would be the number of outlets the machine has
	 */
	for outletIdx := 0; outletIdx < coffeeMachine.Outlets; outletIdx++ {
		wg.Add(1)
		go prepareBeverage(coffeeMachine, beverageChannel, &wg)
	}

	// Only exit when all go routines have finished
	wg.Wait()
}
