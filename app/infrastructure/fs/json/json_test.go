package json_test

import (
	"reflect"
	"testing"

	recipeDomain "github.com/muzfr7/muzfr7-recipe-count-test-2020/app/domain/recipe"
	fsJson "github.com/muzfr7/muzfr7-recipe-count-test-2020/app/infrastructure/fs/json"
)

const (
	testFilePath = "../../../../testdata/test_calculation_fixtures.json"
)

// Test_reader_Read is a unit test for json -> Read() method.
func Test_reader_Read(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    recipeDomain.Data
		wantErr bool
	}{
		{
			name: "Happy path",
			arg:  testFilePath,
			want: recipeDomain.Data{
				Postcode: "10205",
				Recipe:   "Tex-Mex Tilapia",
				Delivery: "Wednesday 1AM - 7PM",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := fsJson.NewReader()

			ch := make(chan recipeDomain.Data)
			go func() {
				defer close(ch)
				ch <- tt.want
			}()

			got, err := r.Read(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("reader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(<-got, <-ch) {
				t.Errorf("reader.Read() = %v, want %v", <-got, <-ch)
			}
		})
	}
}
