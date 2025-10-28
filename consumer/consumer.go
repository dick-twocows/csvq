package consumer

type RowConsumer interface {
	// Accept a row
	// Return true if the consumer can accept the next row or false if not, This allows short circuits, e.g. consuming the header row.
	// Return error if the consumer failed to process the row or nil.
	Accept(Row) (bool, error)
}

type Row interface {
	Index() int
	Values() []string
}

type row struct {
	index  int
	values []string
}

func (self *row) Index() int {
	return self.index
}

func (self *row) Values() []string {
	return self.values
}

func NewRow(index int, values []string) *row {
	return &row{index, values}
}

var EmptyRow = NewRow(0, []string{})
