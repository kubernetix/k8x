package cmd

import (
	"fmt"
	"github.com/kubernetix/k8x/v1/internal/tsx"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(install)
}

var install = &cobra.Command{
	Use:   "install",
	Short: "Install a chart.tsx file into your k8s cluster",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(-1)
		}

		path := args[0]

		code := tsx.Load(path)
		result := tsx.Run(code)

		fmt.Println(result)
	},
}
