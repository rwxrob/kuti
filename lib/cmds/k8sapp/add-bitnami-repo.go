package k8sapp

import (
	"github.com/rwxrob/kuti/lib"
	"github.com/spf13/cobra"
)

var addBitnamiRepo = &cobra.Command{
	Use:   `add-bitnami-repo`,
	Short: `Add Helm bitnami repo and update`,
	RunE: func(_ *cobra.Command, _ []string) error {
		err := lib.Exec(`helm`, `repo`, `add`, `bitnami`,
			`https://charts.bitnami.com/bitnami`)
		if err != nil {
			return err
		}
		return lib.Exec(`helm`, `repo`, `update`, `bitnami`)
	},
}
