package recipe

import (
	"fmt"
	"sort"
	"strings"

	recipeDomain "github.com/muzfr7/muzfr7-recipe-count-test-2020/app/domain/recipe"
)

// Service defines contracts for recipe service.
type Service interface {
	CalculateStats(ch <-chan recipeDomain.Data, filterByPostcodeAndTime recipeDomain.DeliveriesByPostcodeAndTime, filterByWords []string) recipeDomain.ExpectedOutput
}

// service implements Service interface.
type service struct{}

// NewService returns a new instance of service.
func NewService() Service {
	return &service{}
}

// CalculateStats will calculate and return stats as ExpectedOutput.
func (s *service) CalculateStats(ch <-chan recipeDomain.Data, filterByPostcodeAndTime recipeDomain.DeliveriesByPostcodeAndTime, filterByWords []string) recipeDomain.ExpectedOutput {
	var countPerRecipeMap = make(map[string]uint)
	var countPerPostcodeMap = make(map[string]uint)
	var countDeliveriesPerPostcodeMap = make(map[string]uint)
	var filteredRecipeNames []string

	for d := range ch {
		// populate countPerRecipeMap
		if _, ok := countPerRecipeMap[d.Recipe]; !ok {
			countPerRecipeMap[d.Recipe] = 1
		} else {
			countPerRecipeMap[d.Recipe] += 1
		}

		// populate countPerPostcodeMap
		if _, ok := countPerPostcodeMap[d.Postcode]; !ok {
			countPerPostcodeMap[d.Postcode] = 1
		} else {
			countPerPostcodeMap[d.Postcode] += 1
		}

		// populate countDeliveriesPerPostcodeMap
		if d.Postcode == filterByPostcodeAndTime.Postcode && d.InDeliveryTime(filterByPostcodeAndTime) {
			countDeliveriesPerPostcodeMap[d.Postcode] += 1
		}

		// populate filteredRecipeNames
		for _, word := range filterByWords {
			if strings.Contains(d.Recipe, word) && !func() bool {
				for _, filteredRecipeName := range filteredRecipeNames {
					if filteredRecipeName == d.Recipe {
						return true
					}
				}
				return false
			}() {
				filteredRecipeNames = append(filteredRecipeNames, d.Recipe)
				break
			}
		}
	}

	return s.buildExpectedOutput(countPerRecipeMap, countPerPostcodeMap, countDeliveriesPerPostcodeMap, filterByPostcodeAndTime, filteredRecipeNames)
}

// buildExpectedOutput will build final ExpectedOutput.
func (s *service) buildExpectedOutput(countPerRecipeMap, countPerPostcodeMap, countDeliveriesPerPostcodeMap map[string]uint,
	filterByPostcodeAndTime recipeDomain.DeliveriesByPostcodeAndTime, filteredRecipeNames []string) recipeDomain.ExpectedOutput {

	var out recipeDomain.ExpectedOutput

	// set UniqueRecipeCount
	out.UniqueRecipeCount = s.createUniqueRecipeCountFrom(countPerRecipeMap)
	// set CountPerRecipe
	out.CountPerRecipe = s.createSortedCountPerRecipeFrom(countPerRecipeMap)
	// set BusiestPostcode
	out.BusiestPostcode = s.createBusiestPostcodeFrom(countPerPostcodeMap)
	// set CountPerPostcodeAndTime
	out.CountPerPostcodeAndTime = s.createCountPerPostcodeAndTimeFrom(countDeliveriesPerPostcodeMap, filterByPostcodeAndTime)
	// set MatchByName
	sort.Strings(filteredRecipeNames)
	out.MatchByName = filteredRecipeNames

	return out
}

// createUniqueRecipeCountFrom will create UniqueRecipeCount from given countPerRecipeMap.
func (s *service) createUniqueRecipeCountFrom(countPerRecipeMap map[string]uint) uint {
	var out uint

	for _, count := range countPerRecipeMap {
		if count >= 1 {
			out += 1
		}
	}

	return out
}

// createSortedCountPerRecipeFrom will create sorted CountPerRecipe from given countPerRecipeMap.
func (s *service) createSortedCountPerRecipeFrom(countPerRecipeMap map[string]uint) []recipeDomain.CountPerRecipe {
	// create list of recipe names
	var recipeList = make([]string, 0, len(countPerRecipeMap))
	for recipe := range countPerRecipeMap {
		recipeList = append(recipeList, recipe)
	}

	// sort recipe names
	sort.Strings(recipeList)

	var out []recipeDomain.CountPerRecipe

	// create and append CountPerRecipe to out
	for _, recipe := range recipeList {
		out = append(out, recipeDomain.CountPerRecipe{
			Recipe: recipe,
			Count:  countPerRecipeMap[recipe],
		})
	}

	return out
}

// createBusiestPostcodeFrom will create BusiestPostcode from given countPerPostcodeMap
// the BusiestPostcode is the one with most delivered recipes.
func (s *service) createBusiestPostcodeFrom(countPerPostcodeMap map[string]uint) recipeDomain.BusiestPostcode {
	// create list of BusiestPostcode
	var postcodeList []recipeDomain.BusiestPostcode
	for postcode, count := range countPerPostcodeMap {
		postcodeList = append(postcodeList, recipeDomain.BusiestPostcode{
			Postcode:      postcode,
			DeliveryCount: count,
		})
	}

	// sort BusiestPostcode in postcodeList based on DeliveryCount
	sort.Slice(postcodeList, func(i, j int) bool {
		return postcodeList[i].DeliveryCount > postcodeList[j].DeliveryCount
	})

	// since 1st item has highest deliveries
	return postcodeList[0]
}

// createCountPerPostcodeAndTimeFrom will create CountPerPostcodeAndTime from given countDeliveriesPerPostcodeMap.
func (s *service) createCountPerPostcodeAndTimeFrom(countDeliveriesPerPostcodeMap map[string]uint, filterByPostcodeAndTime recipeDomain.DeliveriesByPostcodeAndTime) recipeDomain.CountPerPostcodeAndTime {
	return recipeDomain.CountPerPostcodeAndTime{
		Postcode:      filterByPostcodeAndTime.Postcode,
		FromAM:        fmt.Sprintf("%vAM", filterByPostcodeAndTime.From),
		ToPM:          fmt.Sprintf("%vPM", filterByPostcodeAndTime.To),
		DeliveryCount: countDeliveriesPerPostcodeMap[filterByPostcodeAndTime.Postcode],
	}
}
