package consumer_test

import (
	"fmt"
	"testing"

	"github.com/dick-twocows/csvq/consumer"
	"github.com/dick-twocows/csvq/data"
	"github.com/dick-twocows/csvq/stream"
)

func TestXxx(t *testing.T) {
	f, _ := data.FS.Open("people-10.csv")

	defer func() {
		f.Close()
	}()

	control := stream.NewControl()

	csvSource := consumer.NewCSVSource(control, f)
	csvSource.Start()

	forEach := stream.NewForEach(control, csvSource, stream.NewStdOutForEachConsumer)
	fmt.Printf("err [%v]\n", forEach.Start())
}

func TestHeaders(t *testing.T) {
	f, _ := data.FS.Open("people-10.csv")

	defer func() {
		f.Close()
	}()

	control := stream.NewControl()

	defer func() {
		control.Stop()
	}()

	csvSource := consumer.NewCSVSource(control, f)
	csvSource.Start()

	p := func(_ stream.Control, r consumer.Row) error {
		fmt.Printf("peek r [%v]\n", r)
		return nil
	}

	csvHeaders := consumer.NewHeadersIntermediate(control, csvSource)
	csvHeaders.Start()

	peek := stream.NewForEachIntermediate(control, csvHeaders, p)
	peek.Start()

	forEach := stream.NewForEach(control, csvHeaders, stream.NewStdOutForEachConsumer)
	forEach.Start()

}

func TestCSVRange(t *testing.T) {
	f, _ := data.FS.Open("people-10.csv")

	defer func() {
		f.Close()
	}()

	control := stream.NewControl()

	defer func() {
		control.Stop()
	}()

	csvSource := consumer.NewCSVSource(control, f)
	csvSource.Start()

	p := func(_ stream.Control, r consumer.Row) error {
		fmt.Printf("peek r [%v]\n", r)
		return nil
	}

	// csvHeaders := consumer.NewHeadersIntermediate(control, csvSource)
	// csvHeaders.Start()

	peek := stream.NewForEachIntermediate(control, csvSource, p)
	peek.Start()

	r := consumer.NewCSVRangeIntermediate(control, peek, 4, 6)
	r.Start()

	forEach := stream.NewForEach(control, r, stream.NewStdOutForEachConsumer)
	forEach.Start()

}
