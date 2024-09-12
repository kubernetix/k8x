package cmd

import (
	"github.com/kubernetix/k8x/v1/internal/k8s"
	"github.com/kubernetix/k8x/v1/internal/tsx"
	"github.com/spf13/cobra"
	"os"
)

var Verbose bool

func init() {
	install.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Enable debug output")
	rootCmd.AddCommand(install)
}

var install = &cobra.Command{
	Use:   "install",
	Short: "Install a chart.tsx file into your k8s cluster",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			os.Exit(-1)
		}

		path := args[0]

		code := tsx.Load(path, Verbose)
		_ = tsx.Run(code)

		k8s.CreateNamespae("default")
	},
}
