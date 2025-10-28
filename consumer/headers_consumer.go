package consumer

import (
	"errors"

	"github.com/dick-twocows/csvq/stream"
)

var CSVHeadersRowNE1 = errors.New("CSV Headers row != 1")

func NewHeadersIntermediate(control stream.Control, in stream.Source[Row]) stream.Filter[Row] {

	count := 0

	predicate := func(_ stream.Control, _ Row) (bool, error) {
		count++

		if count == 1 {
			return true, nil
		}

		return false, CSVHeadersRowNE1
	}

	return stream.NewFilterIntermediate(control, in, predicate)
}

// func csvHeadersPretty() (string, error) {
// 	t := table.NewWriter()

// 	t.SetAutoIndex(true)
// 	t.AppendHeader(table.Row{"Header"})

// 	for _, v := range self.headers.Values() {
// 		t.AppendRow([]interface{}{v})
// 	}

// 	return t.Render(), nil
// }
