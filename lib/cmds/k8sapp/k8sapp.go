package k8sapp

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(removeSecrets)
}

var Cmd = &cobra.Command{
	Use: `k8sapp`,
}
