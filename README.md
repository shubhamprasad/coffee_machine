Coffee Machine.

Tests cases are present in test_inputs directory.
The easiest way to test this would be execute the binary "coffee_machine" by passing the input file name as a flag

./coffee_machine -file test_inputs/input.json

Other functions like GetLowRunningIngredients and AddIngredients are commented out to keep the output simple.
If you want to test them uncomment it in main.go and build the project.

Install Go in your environment

go build -o coffee_machine main.go
./coffee_machine -file test_inputs/input.json

Tried to make it very simple with minimal requirements and assumptions.

To test whether all N outlets are working in parallel, you should be getting a different output most of the times for
the same input file.

The number of outlets would be the number of goroutines (lightweight threads) spawned.
Each worker will pick a beverage from the channel and prepare it.

Since multiple workers will access the coffee machine's ingredients, mutex's are used to avoid race conditions in
the critical section.



