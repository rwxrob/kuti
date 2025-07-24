package k8sapp

import (
	"log"

	"github.com/spf13/cobra"
)

var removeSecrets = &cobra.Command{
	Use:  `remove-secrets`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(_ *cobra.Command, files []string) error {
		log.Print(`would remove secrets from following files`)
		for _, name := range files {
			log.Println(name)
		}
		return nil
	},
}
