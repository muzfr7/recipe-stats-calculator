package recipe_test

import (
	"testing"

	recipeDomain "github.com/muzfr7/muzfr7-recipe-count-test-2020/app/domain/recipe"
)

// TestData_InDeliveryTime is a unit test for recipe -> InDeliveryTime() method.
func TestData_InDeliveryTime(t *testing.T) {
	tests := []struct {
		name string
		data *recipeDomain.Data
		arg  recipeDomain.DeliveriesByPostcodeAndTime
		want bool
	}{
		{
			name: "Happy path",
			data: &recipeDomain.Data{
				Postcode: "10120",
				Recipe:   "Creamy Dill Chicken",
				Delivery: "Thursday 11AM - 2PM",
			},
			arg: recipeDomain.DeliveriesByPostcodeAndTime{
				Postcode: "10120",
				From:     10,
				To:       3,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.data.InDeliveryTime(tt.arg); got != tt.want {
				t.Errorf("Data.InDeliveryTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
