package cmd

import "github.com/spf13/cobra"

var CSVFile string

var RootCmd = &cobra.Command{
	Use:   "csvq [text]",
	Short: "A simple CSV query tool",
	Args:  cobra.ArbitraryArgs,
}

func init() {
	RootCmd.PersistentFlags().StringVar(&CSVFile, "file", "", "CSV file")
}
