package consumer

import (
	"slices"

	"github.com/jedib0t/go-pretty/v6/table"
)

type RowsConsumer struct {
	indexes []int
	count   int
	rows    [][]string
}

func (self *RowsConsumer) Accept(row []string) (bool, error) {
	self.count++
	if slices.Contains(self.indexes, self.count) {
		self.rows = append(self.rows, row)
	}

	return true, nil
}

func (self *RowsConsumer) Pretty() (string, error) {
	t := table.NewWriter()

	t.SetAutoIndex(true)
	t.AppendHeader(table.Row{"Header"})

	// for _, v := range self.headers {
	// 	t.AppendRow([]interface{}{v})
	// }

	return t.Render(), nil
}

func NewRowsConsumer(indexes []int) *RowsConsumer {
	return &RowsConsumer{indexes, 0, [][]string{}}
}
