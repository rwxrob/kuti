package cmds

import (
	"github.com/rwxrob/kuti/lib/cmds/k8sapp"
	"github.com/spf13/cobra"
)

func init() {
	Kuti.AddCommand(k8sapp.Cmd)
}

var Kuti = &cobra.Command{
	Use:   `kuti`,
	Short: `Collection of Kubernetes utilities`,
}
