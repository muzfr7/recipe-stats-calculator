package recipe_test

import (
	"reflect"
	"testing"

	recipeDomain "github.com/muzfr7/muzfr7-recipe-count-test-2020/app/domain/recipe"
	recipeUsecase "github.com/muzfr7/muzfr7-recipe-count-test-2020/app/usecases/recipe"
)

// Test_service_CalculateStats is a unit test for recipe -> CalculateStats() method.
func Test_service_CalculateStats(t *testing.T) {
	type args struct {
		ch                      recipeDomain.Data
		filterByPostcodeAndTime recipeDomain.DeliveriesByPostcodeAndTime
		filterByWords           []string
	}

	tests := []struct {
		name string
		args args
		want recipeDomain.ExpectedOutput
	}{
		{
			name: "Happy path",
			args: args{
				ch: recipeDomain.Data{
					Postcode: "10205",
					Recipe:   "Tex-Mex Tilapia",
					Delivery: "Wednesday 11AM - 3PM",
				},
				filterByPostcodeAndTime: recipeDomain.DeliveriesByPostcodeAndTime{
					Postcode: "10205",
					From:     10,
					To:       3,
				},
				filterByWords: []string{"Tilapia"},
			},
			want: recipeDomain.ExpectedOutput{
				UniqueRecipeCount: 1,
				CountPerRecipe: []recipeDomain.CountPerRecipe{
					{
						Recipe: "Tex-Mex Tilapia",
						Count:  1,
					},
				},
				BusiestPostcode: recipeDomain.BusiestPostcode{
					Postcode:      "10205",
					DeliveryCount: 1,
				},
				CountPerPostcodeAndTime: recipeDomain.CountPerPostcodeAndTime{
					Postcode:      "10205",
					FromAM:        "10AM",
					ToPM:          "3PM",
					DeliveryCount: 1,
				},
				MatchByName: []string{"Tex-Mex Tilapia"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := recipeUsecase.NewService()

			ch := make(chan recipeDomain.Data)
			go func() {
				defer close(ch)
				ch <- tt.args.ch
			}()

			if got := svc.CalculateStats(ch, tt.args.filterByPostcodeAndTime, tt.args.filterByWords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CalculateStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
