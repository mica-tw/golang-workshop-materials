package cmd

import (
	"concat/internal/service"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "concat",
	Short: "Merge files concurrently and convert to uppercase",
	RunE: func(cmd *cobra.Command, args []string) error {
		inputDir, err := cmd.Flags().GetString("dir")
		if err != nil {
			return err
		}
		outputFile, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}
		numWorkers, err := cmd.Flags().GetInt("workers")
		if err != nil {
			return err
		}

		svc := service.NewDirService(inputDir, outputFile, numWorkers)

		return svc.ProcessDir()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("dir", "d", ".", "directory containing files to merge")
	rootCmd.Flags().StringP("output", "o", "output.txt", "output file name")
	rootCmd.Flags().IntP("workers", "w", 5, "number of workers to use")
}
