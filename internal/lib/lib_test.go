package lib_test

import (
	"fmt"

	"github.com/rwxrob/kuti/internal/lib"
)

func Example_YQ() {

	yaml := `
some: thing
version: v1.0.0
`

	out, err := lib.YQ(`.version`, yaml)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)

	// Output:
	// v1.0.0
}
