package consumer

import (
	"encoding/csv"
	"io/fs"
	"os"

	"github.com/dick-twocows/csvq/stream"
)

const (
	FirstRow int = iota
	UserDefined
)

type CSVSourceContext struct {
	file          string
	headersSource int
	headers       Row
}

type CSVSource interface {
	stream.Source[Row]
}

type csvSource struct {
	context        CSVSourceContext
	control        stream.Control
	supplierSource stream.Source[Row]
}

func (self *csvSource) Control() stream.Control {
	return self.supplierSource.Control()
}

func (self *csvSource) Out() <-chan Row {
	return self.supplierSource.Out()
}

func (self *csvSource) Start() error {
	// started := false

	f, err := os.Open(self.context.file)
	if err != nil {
		return err
	}

	defer func() {
		f.Close()
	}()

	r := csv.NewReader(f)

	index := 0

	supplier := func(_ stream.Control) (Row, bool, error) {
		index++

		record, err := r.Read()
		if err != nil {
			return EmptyRow, false, err
		}

		return NewRow(index, record), true, nil
	}

	self.supplierSource = stream.NewSupplierSource(self.control, supplier)
	self.supplierSource.Start()

	return nil
}

func NewCSVSource(control stream.Control, f fs.File) stream.Source[Row] {

	r := csv.NewReader(f)

	index := 0

	supplier := func(_ stream.Control) (Row, bool, error) {
		index++

		record, err := r.Read()
		if err != nil {
			return EmptyRow, false, err
		}

		return NewRow(index, record), true, nil
	}

	return stream.NewSupplierSource(control, supplier)
}
