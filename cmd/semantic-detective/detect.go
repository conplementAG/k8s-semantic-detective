package main

import (
	"github.com/conplementAG/k8s-semantic-detective/pkg/detective"
	"github.com/spf13/cobra"
)

func createDetectCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "detect",
		Short: "Start the detective",
		Long: `
Use this command to start the detective exporter.
		`,
		Run: func(cmd *cobra.Command, args []string) {
			detective.Detect()
		},
	}

	return command
}
