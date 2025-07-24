package cmds

import (
	"fmt"

	"github.com/rwxrob/kuti/lib/cmds/k8sapp"
	"github.com/spf13/cobra"
)

func init() {
	Kuti.AddCommand(say, k8sapp.Cmd)
}

var Kuti = &cobra.Command{
	Use: `kuti`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("would run kuti")
	},
}
