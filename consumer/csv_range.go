package consumer

import (
	"errors"

	"github.com/dick-twocows/csvq/stream"
)

var CSVRangeToExceeded = errors.New("CSV range to exceeded")

func NewCSVRangeIntermediate(control stream.Control, in stream.Source[Row], from int, to int) stream.Filter[Row] {

	predicate := func(_ stream.Control, row Row) (bool, error) {

		if row.Index() < from {
			return false, nil
		}

		if row.Index() <= to {
			return true, nil
		}

		return false, CSVRangeToExceeded
	}

	return stream.NewFilterIntermediate(control, in, predicate)
}
