package recipe_test

import (
	"testing"

	recipeDomain "github.com/muzfr7/muzfr7-recipe-count-test-2020/app/domain/recipe"
)

// TestExpectedOutput_Jsonify is a unit test for recipe -> Jsonify() method.
func TestExpectedOutput_Jsonify(t *testing.T) {
	tests := []struct {
		name    string
		data    *recipeDomain.ExpectedOutput
		want    string
		wantErr bool
	}{
		{
			name: "Happy path",
			data: &recipeDomain.ExpectedOutput{
				UniqueRecipeCount: 3,
				CountPerRecipe: []recipeDomain.CountPerRecipe{
					{
						Recipe: "Tex-Mex Tilapia",
						Count:  1,
					}, {
						Recipe: "Grilled Cheese and Veggie Jumble",
						Count:  1,
					}, {
						Recipe: "Creamy Dill Chicken",
						Count:  1,
					},
				},
				BusiestPostcode: recipeDomain.BusiestPostcode{
					Postcode:      "10205",
					DeliveryCount: 2,
				},
				CountPerPostcodeAndTime: recipeDomain.CountPerPostcodeAndTime{
					Postcode:      "10120",
					FromAM:        "10AM",
					ToPM:          "3PM",
					DeliveryCount: 1,
				},
				MatchByName: []string{"Grilled Cheese and Veggie Jumble"},
			},
			want: `{
	"unique_recipe_count": 3,
	"count_per_recipe": [
		{
			"recipe": "Tex-Mex Tilapia",
			"count": 1
		},
		{
			"recipe": "Grilled Cheese and Veggie Jumble",
			"count": 1
		},
		{
			"recipe": "Creamy Dill Chicken",
			"count": 1
		}
	],
	"busiest_postcode": {
		"postcode": "10205",
		"delivery_count": 2
	},
	"count_per_postcode_and_time": {
		"postcode": "10120",
		"from": "10AM",
		"to": "3PM",
		"delivery_count": 1
	},
	"match_by_name": [
		"Grilled Cheese and Veggie Jumble"
	]
}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.data.Jsonify()
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpectedOutput.Jsonify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("ExpectedOutput.Jsonify() = %v, want %v", got, tt.want)
			}
		})
	}
}
