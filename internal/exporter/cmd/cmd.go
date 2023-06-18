// // Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// // Licensed under the MIT License (the "License");
// // you may not use this file except in compliance with the License.

// package cmd

// import (
// 	"flag"
// 	"io"
// 	"os"

// 	"github.com/kubecub/github-label-syncer/internal/utils/templates"
// 	"github.com/spf13/cobra"
// 	"github.com/spf13/viper"
// 	"k8s.io/kubectl/pkg/cmd/completion"
// 	"k8s.io/kubectl/pkg/cmd/options"
// 	"k8s.io/kubectl/pkg/cmd/set"
// 	"k8s.io/kubectl/pkg/cmd/version"
// 	"sigs.k8s.io/kustomize/kustomize/v4/commands/openapi/info"
// )

// // NewDefaultIAMCtlCommand creates the `iamctl` command with default arguments.
// func NewDefaultIAMCtlCommand() *cobra.Command {
// 	return NewEeporterCtlCommand(os.Stdin, os.Stdout, os.Stderr)
// }

// var LongDesc = templates.LongDesc(`
// iamctl controls the iam platform, is the client side tool for iam platform.

// Find more information at:
// 	https://github.com/marmotedu/iam/blob/master/docs/guide/en-US/cmd/iamctl/iamctl.md`)

// // NewEeporterCtlCommand returns new initialized instance of 'exporter' root command.
// func NewEeporterCtlCommand(in io.Reader, out, err io.Writer) *cobra.Command {
// 	cmds := &cobra.Command{
// 		Use:   "exporter",
// 		Short: "exporter exporting labels",
// 		Long:  LongDesc,
// 		Run:   runHelp,
// 		// Hook before and after Run initialize and write profiles to disk,
// 		// respectively.
// 		PersistentPreRunE: func(*cobra.Command, []string) error {
// 			return initProfiling()
// 		},
// 		PersistentPostRunE: func(*cobra.Command, []string) error {
// 			return flushProfiling()
// 		},
// 	}

// 	flags := cmds.PersistentFlags()
// 	flags.SetNormalizeFunc(cliflag.WarnWordSepNormalizeFunc) // Warn for "_" flags

// 	// Normalize all flags that are coming from other packages or pre-configurations
// 	// a.k.a. change all "_" to "-". e.g. glog package
// 	flags.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)

// 	addProfilingFlags(flags)

// 	iamConfigFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag().WithDeprecatedSecretFlag()
// 	iamConfigFlags.AddFlags(flags)
// 	matchVersionIAMConfigFlags := cmdutil.NewMatchVersionFlags(iamConfigFlags)
// 	matchVersionIAMConfigFlags.AddFlags(cmds.PersistentFlags())

// 	_ = viper.BindPFlags(cmds.PersistentFlags())
// 	cobra.OnInitialize(func() {
// 		genericapiserver.LoadConfig(viper.GetString(genericclioptions.FlagIAMConfig), "iamctl")
// 	})
// 	cmds.PersistentFlags().AddGoFlagSet(flag.CommandLine)

// 	f := cmdutil.NewFactory(matchVersionIAMConfigFlags)

// 	// From this point and forward we get warnings on flags that contain "_" separators
// 	cmds.SetGlobalNormalizationFunc(cliflag.WarnWordSepNormalizeFunc)

// 	ioStreams := genericclioptions.IOStreams{In: in, Out: out, ErrOut: err}

// 	groups := templates.CommandGroups{
// 		{
// 			Message: "Basic Commands:",
// 			Commands: []*cobra.Command{
// 				info.NewCmdInfo(f, ioStreams),
// 				color.NewCmdColor(f, ioStreams),
// 				new.NewCmdNew(f, ioStreams),
// 				jwt.NewCmdJWT(f, ioStreams),
// 			},
// 		},
// 		{
// 			Message: "Identity and Access Management Commands:",
// 			Commands: []*cobra.Command{
// 				user.NewCmdUser(f, ioStreams),
// 				secret.NewCmdSecret(f, ioStreams),
// 				policy.NewCmdPolicy(f, ioStreams),
// 			},
// 		},
// 		{
// 			Message: "Troubleshooting and Debugging Commands:",
// 			Commands: []*cobra.Command{
// 				validate.NewCmdValidate(f, ioStreams),
// 			},
// 		},
// 		{
// 			Message: "Settings Commands:",
// 			Commands: []*cobra.Command{
// 				set.NewCmdSet(f, ioStreams),
// 				completion.NewCmdCompletion(ioStreams.Out, ""),
// 			},
// 		},
// 	}
// 	groups.Add(cmds)

// 	filters := []string{"options"}
// 	templates.ActsAsRootCommand(cmds, filters, groups...)

// 	cmds.AddCommand(version.NewCmdVersion(f, ioStreams))
// 	cmds.AddCommand(options.NewCmdOptions(ioStreams.Out))

// 	return cmds
// }

// func runHelp(cmd *cobra.Command, args []string) {
// 	_ = cmd.Help()
// }
