package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var say = &cobra.Command{
	Use: `say`,
	Run: func(_ *cobra.Command, args []string) {
		fmt.Print(`Saying`)
		fmt.Println(args)
	},
}
