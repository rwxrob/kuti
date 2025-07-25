package k8sapp

import (
	"github.com/spf13/cobra"
)

var pullDefaultValues = &cobra.Command{
	Use:  `pull-default-values`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		//lib.Exec()
		//	helm pull oci://registry-1.docker.io/grafana/grafana --version <VERSION>
		return nil
	},
}
