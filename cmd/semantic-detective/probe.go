package main

import (
	"github.com/conplementAG/k8s-semantic-detective/internal/externalprobe"
	"github.com/spf13/cobra"
)

func createProbeCommand() *cobra.Command {
	var namespace = ""
	var command = &cobra.Command{
		Use:   "probe",
		Short: "Checks if the k8s management api is functional",
		Long: `
Use this command adsfasdf to check if the k8s management api is functional. The probe will default to the 
management api of cluster its running within.
		`,
		Run: func(cmd *cobra.Command, args []string) {
			externalprobe.Probe(namespace)
		},
	}

	command.PersistentFlags().StringVar(&namespace, "namespace", "","Set namespace where probe artifacts should be created in")
	command.MarkPersistentFlagRequired("namespace")


	return command
}