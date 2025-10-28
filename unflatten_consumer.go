package main

import (
	"github.com/jedib0t/go-pretty/v6/list"
)

type UnflattenConsumer struct {
	columns []int
	counts  map[string]int
	nodes   map[string]UnflattenConsumer
}

func (self *UnflattenConsumer) Accept(row []string) error {

	// Increment count for cell.
	self.counts[row[self.columns[0]-1]] = self.counts[row[self.columns[0]-1]] + 1

	// Check if we are a leaf node.
	if len(self.columns) == 1 {
		return nil
	}

	// Call accept on cell node.
	var n UnflattenConsumer
	var ok bool
	n, ok = self.nodes[row[self.columns[0]-1]]
	if !ok {
		n = *NewUnflattenConsumer(self.columns[1:])
		self.nodes[row[self.columns[0]-1]] = n
	}
	n.Accept(row)

	return nil
}

func (self *UnflattenConsumer) Pretty() (string, error) {
	l := list.NewWriter()
	l.SetStyle(list.StyleConnectedRounded)

	self.pretty(l)

	return l.Render(), nil
}

func (self *UnflattenConsumer) pretty(l list.Writer) error {

	if len(self.columns) == 1 {
		// fmt.Printf("%sLEAF [%v]", indent, len(self.counts))
		for k := range self.counts {
			l.AppendItem(k)
		}
		return nil
	}

	for k := range self.counts {
		l.AppendItem(k)
		l.Indent()
		n := self.nodes[k]
		n.pretty(l)
		l.UnIndent()
	}
	return nil
}

func NewUnflattenConsumer(columns []int) *UnflattenConsumer {
	return &UnflattenConsumer{columns, map[string]int{}, map[string]UnflattenConsumer{}}
}
