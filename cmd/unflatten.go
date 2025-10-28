package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var unflattenColumns []int

var unflattenCmd = &cobra.Command{
	Use:   "unflatten [description]",
	Short: "Unflatten a CSB file",
	Args:  cobra.MinimumNArgs(1),
	Run:   unflattenTask,
}

func unflattenTask(cmd *cobra.Command, args []string) {
	fmt.Printf("[%v]\n\n[%v]", cmd, args)
}

func init() {
	unflattenCmd.Flags().IntSliceVar(&unflattenColumns, "columns", []int{}, "Columns to unflatten")

	RootCmd.AddCommand(unflattenCmd)
}
