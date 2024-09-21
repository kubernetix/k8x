package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/kubernetix/k8x/v1/internal/ts"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

var Json bool

func init() {
	render.PersistentFlags().BoolVarP(&Json, "json", "j", false, "Render json")
	rootCmd.AddCommand(render)
}

var render = &cobra.Command{
	Use:   "render",
	Short: "Render a chart file as yaml or json",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			os.Exit(-1)
		}

		path := args[0]

		code := ts.Load(path, Verbose)
		export := ts.Run(code)

		if Json {
			jsn, _ := json.MarshalIndent(export, "", "  ")
			fmt.Println(string(jsn))
		} else {
			yml, _ := yaml.Marshal(export)
			fmt.Println(string(yml))
		}

	},
}
