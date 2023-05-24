package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/kubecub/github-label-syncer/internal/utils/templates"
	"github.com/spf13/cobra"
)

// NewDefaultIAMCtlCommand creates the `iamctl` command with default arguments.
func NewDefaultIAMCtlCommand() *cobra.Command {
	return NewEeporterCtlCommand(os.Stdin, os.Stdout, os.Stderr)
}

// NewEeporterCtlCommand returns new initialized instance of 'exporter' root command.
func NewEeporterCtlCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "exporter",
		Short: "exporter exporting labels",
		Long: templates.LongDesc(`
		iamctl controls the iam platform, is the client side tool for iam platform.

		Find more information at:
			https://github.com/marmotedu/iam/blob/master/docs/guide/en-US/cmd/iamctl/iamctl.md`),
		Run: runHelp,
		// Hook before and after Run initialize and write profiles to disk,
		// respectively.
		PersistentPreRunE: func(*cobra.Command, []string) error {
			return initProfiling()
		},
		PersistentPostRunE: func(*cobra.Command, []string) error {
			return flushProfiling()
		},
	}

	return cmds
}

func runHelp() {
	fmt.Println("run help")
}