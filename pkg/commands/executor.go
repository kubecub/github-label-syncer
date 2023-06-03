package commands

import "github.com/spf13/cobra"

type BuildInfo struct {
	GoVersion string `json:"goVersion"`
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	Date      string `json:"date"`
}

type Executor struct {
	rootCmd    *cobra.Command
	runCmd     *cobra.Command
	lintersCmd *cobra.Command

	exitCode  int
	buildInfo BuildInfo
}
