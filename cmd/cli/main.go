package main

import (
	"fmt"
	"os"

	recipeDomain "github.com/muzfr7/muzfr7-recipe-count-test-2020/app/domain/recipe"
	fsJson "github.com/muzfr7/muzfr7-recipe-count-test-2020/app/infrastructure/fs/json"
	recipeUsecase "github.com/muzfr7/muzfr7-recipe-count-test-2020/app/usecases/recipe"
)

const recipeDataFile string = "./resources/fixtures/hf_test_calculation_fixtures.json"

func main() {
	// instantiate json reader
	jsonReader := fsJson.NewReader()

	// instantiate recipe service
	recipeSVC := recipeUsecase.NewService()

	// read each recipe data in ch
	ch, err := jsonReader.Read(recipeDataFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// to count the number of deliveries to postcode 10120 that lie within the delivery time between 10AM and 3PM
	filterByPostcodeAndTime := recipeDomain.DeliveriesByPostcodeAndTime{
		Postcode: "10120",
		From:     10,
		To:       3,
	}

	// to list the recipe names that contain one of following words in their name
	filterByWords := []string{"Potato", "Veggie", "Mushroom"}

	// call CalculateStats method from recipeSVC
	output := recipeSVC.CalculateStats(ch, filterByPostcodeAndTime, filterByWords)
	jsonOutput, err := output.Jsonify()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%v\n", jsonOutput)
}
