package main

import (
	"coffee_machine/models"
	"coffee_machine/usecases"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

func main()  {

	inputFileName := flag.String("file", "input.json", "")

	flag.Parse()

	inputFile, err := ioutil.ReadFile(*inputFileName)
	if err != nil {
		fmt.Println("Unable to read input file, error: ", err)
		return
	}

	inputStruct := models.InputStruct{}

	// Map json in input file to a struct
	err = json.Unmarshal(inputFile, &inputStruct)

	if err != nil {
		fmt.Println("Unable to map json to struct, error: ", err)
	}

	coffeeMachine, beverageList := usecases.InitializeCoffeeMachine(inputStruct)

	usecases.PrepareBeverages(&coffeeMachine, beverageList)

	//Use this function to check low running ingredients. Pass the desired threshold
	//coffeeMachine.GetLowRunningIngredients(60)

	// Use this function to add ingredients.
	//coffeeMachine.AddIngredients("hot_water", 50)
	//fmt.Println(coffeeMachine.Ingredients)

}
