package json

import (
	"encoding/json"
	"fmt"
	"os"

	recipeDomain "github.com/muzfr7/muzfr7-recipe-count-test-2020/app/domain/recipe"
)

// The number of elements that can be sent to the channel without the blocking.
const bufferSize = 100

// Reader defines contracts for json reader.
type Reader interface {
	Read(filePath string) (<-chan recipeDomain.Data, error)
}

// reader implements Reader interface.
type reader struct{}

// NewReader returns a new instance of reader.
func NewReader() Reader {
	return &reader{}
}

// Read will read and return each object from given JSON file.
func (r *reader) Read(filePath string) (<-chan recipeDomain.Data, error) {
	// open JSON file
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// read and decode JSON values from file
	d := json.NewDecoder(f)

	// read open bracket
	if _, err := d.Token(); err != nil {
		return nil, err
	}

	ch := make(chan recipeDomain.Data, bufferSize)

	go func() {
		defer close(ch)

		// while the array contains values
		for d.More() {
			var data recipeDomain.Data

			// decode an array value (RecipeData)
			if err := d.Decode(&data); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}

			ch <- data
		}

		// read closing bracket
		if _, err := d.Token(); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		r.closeFile(f)
	}()

	return ch, nil
}

// closeFile will close given file.
func (r *reader) closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
