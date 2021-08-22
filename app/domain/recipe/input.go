package recipe

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Data represents a single recipe data obj in hf_test_calculation_fixtures.json file.
type Data struct {
	Postcode string `json:"postcode"`
	Recipe   string `json:"recipe"`
	Delivery string `json:"delivery"`
}

// InDeliveryTime will check given time in Data.Delivery
// d.Delivery has following format: {weekday} {h}AM - {h}PM.
func (d *Data) InDeliveryTime(filterByPostcodeAndTime DeliveriesByPostcodeAndTime) bool {
	regex := regexp.MustCompile(`(\d{0,2})AM\s-\s(\d{0,2})PM`)
	matches := regex.FindStringSubmatch(d.Delivery)

	from, err := strconv.Atoi(matches[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	to, err := strconv.Atoi(matches[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	return uint(from) >= filterByPostcodeAndTime.From && uint(to) <= filterByPostcodeAndTime.To
}
