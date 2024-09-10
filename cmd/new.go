package cmd

import (
	"bytes"
	"fmt"
	spinner2 "github.com/kubernetix/k8x/v1/internal/spinner"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"path"
	"text/template"
	"time"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

var chartTsx = `/// <reference types="@kubernetix/types" />

export default () => (
  <chart name={""} version={""} >
  </chart>
)
`

var packageJson = `{
  "name": "{{.packageName}}",
  "private": true,
  "version": "0.0.0",
  "dependencies": {
    "@kubernetix/types": "0.0.1"
  },
  "chart": {
    "appVersion": "1.0.0",
    "kubeVersion": "1.31",
    "type": "application",
    "keywords": [
      "cms",
      "wordpress",
      "author"
    ],
    "home": "https://github.com/kubernetix/charts/wordpress",
    "maintainers": [
      "Niklas Hanft"
    ],
    "icon": null,
    "deprecated": false,
    "annotations": []
  }
}
`

var newCmd = &cobra.Command{
	Use:     "new",
	Short:   "Initialize a k8x chart. (chart.tsx, package.json)",
	Example: "k8x new wordpress",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(-1)
		}

		chartPath := args[0]

		t, _ := template.New("text").Parse(packageJson)

		pkgjson := bytes.Buffer{}
		err := t.Execute(&pkgjson, map[string]interface{}{"packageName": chartPath})

		if err != nil {
			log.Error().Msg(err.Error())
			os.Exit(-1)
		}

		err = os.Mkdir(chartPath, 0666)

		if err != nil {
			log.Error().Msg(err.Error())
			os.Exit(-1)
		}

		err = os.WriteFile(path.Join(chartPath, "package.json"), pkgjson.Bytes(), 0666)

		if err != nil {
			log.Error().Msg(err.Error())
			os.Exit(-1)
		}

		err = os.WriteFile(path.Join(chartPath, "chart.tsx"), []byte(chartTsx), 0666)

		if err != nil {
			log.Error().Err(err).Msg("")
			os.Exit(-1)
		}

		spinner := spinner2.NewSpinner()

		for i := 0; i < 50; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("\033[2K")
			fmt.Println()
			fmt.Printf("\033[1A")
			fmt.Printf("%s Initializing chart....", spinner.String())
		}
		fmt.Println()
		log.Info().Msg("Success!")
	},
}
