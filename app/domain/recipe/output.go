package recipe

import "encoding/json"

// CountPerRecipe represents a unique recipe with the number of its occurrences in ExpectedOutput.
type CountPerRecipe struct {
	Recipe string `json:"recipe"`
	Count  uint   `json:"count"`
}

// BusiestPostcode represents the postcode with most delivered recipes in ExpectedOutput.
type BusiestPostcode struct {
	Postcode      string `json:"postcode"`
	DeliveryCount uint   `json:"delivery_count"`
}

// CountPerPostcodeAndTime represents delivery count to a postcode at certain time in ExpectedOutput.
type CountPerPostcodeAndTime struct {
	Postcode      string `json:"postcode"`
	FromAM        string `json:"from"`
	ToPM          string `json:"to"`
	DeliveryCount uint   `json:"delivery_count"`
}

// ExpectedOutput represents expected output with required stats.
type ExpectedOutput struct {
	UniqueRecipeCount       uint                    `json:"unique_recipe_count"`
	CountPerRecipe          []CountPerRecipe        `json:"count_per_recipe"`
	BusiestPostcode         BusiestPostcode         `json:"busiest_postcode"`
	CountPerPostcodeAndTime CountPerPostcodeAndTime `json:"count_per_postcode_and_time"`
	MatchByName             []string                `json:"match_by_name"`
}

// Jsonify will convert ExpectedOutput to json string.
func (eo *ExpectedOutput) Jsonify() (string, error) {
	b, err := json.MarshalIndent(eo, "", "\t")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
