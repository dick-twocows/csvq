package cmd

import (
	"fmt"
	"log"

	"github.com/dick-twocows/csvq/consumer"
	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use: "describe [description]",
	// Short: "Describe",
}

var describeHeadersCmd = &cobra.Command{
	Use: "headers [description]",
	// Short: "Describe headers",
	Run: describeHeadersTask,
}

var describeRowCmd = &cobra.Command{
	Use: "row [description]",
	// Short: "Describe a row",
	Run: describeRowTask,
}

func describeHeadersTask(cmd *cobra.Command, args []string) {
	c := consumer.NewHeadersConsumer()
	if err := consumer.FileConsumer(CSVFile, c); err != nil {
		log.Fatal(err)
	}
	res, err := c.Pretty()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", res)
}

func describeRowTask(cmd *cobra.Command, args []string) {
	c := consumer.NewHeadersConsumer()
	if err := consumer.FileConsumer(CSVFile, c); err != nil {
		log.Fatal(err)
	}
	res, err := c.Pretty()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", res)
}

func init() {
	describeCmd.AddCommand(describeHeadersCmd)
	RootCmd.AddCommand(describeCmd)
}
