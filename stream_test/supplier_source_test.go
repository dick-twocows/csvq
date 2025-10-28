package stream_test

import (
	"fmt"
	"testing"

	"github.com/dick-twocows/csvq/stream"
)

func TestSupplierSource1(t *testing.T) {

	control := stream.NewControl()

	count := 0
	supplier := func(_ stream.Control) (int, bool, error) {
		count++
		return count, count != 10, nil
	}

	source := stream.NewSupplierSource(control, supplier)
	source.Start()

	consumer := func(_ stream.Control, t int) error {
		fmt.Printf("consumer t [%v]\n", t)
		return nil
	}

	forEach := stream.NewForEach(control, source, consumer)
	fmt.Printf("err [%v]\n", forEach.Start())
	fmt.Printf("err [%v]\n", forEach.Start())

}
