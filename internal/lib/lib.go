package lib

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Exec checks for existence of first argument as an executable on the
// system and then runs it with [exec.Command.Run]  exiting in a way that
// is supported across all architectures that Go supports. The [os.Stdin],
// [os.Stdout], and [os.Stderr] are connected directly to that of the calling
// program. Sometimes this is insufficient and the UNIX-specific SysExec
// is preferred.
func Exec(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing name of executable")
	}
	path, err := exec.LookPath(args[0])
	if err != nil {
		return err
	}
	cmd := exec.Command(path, args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ExecOut returns the standard output of the executed command as
// a string. Errors are logged but not returned.
func ExecOut(args ...string) string {
	if len(args) == 0 {
		log.Println("missing name of executable")
		return ""
	}
	path, err := exec.LookPath(args[0])
	if err != nil {
		log.Println(err)
		return ""
	}
	out, err := exec.Command(path, args[1:]...).Output()

	if err != nil {
		log.Println(err)
	}
	return string(out)
}

func YQ(query, yaml string) (string, error) {
	path, err := exec.LookPath(`yq`)
	if err != nil {
		return "", err
	}
	in := strings.NewReader(yaml)
	out := new(bytes.Buffer)
	cmd := exec.Command(path, query)
	cmd.Stdout = out
	cmd.Stdin = in
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return out.String(), err
}
