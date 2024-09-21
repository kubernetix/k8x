package ts

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

// Loads and transpiles tsx files
func Load(path string, debug bool) string {

	options := api.BuildOptions{
		Loader: map[string]api.Loader{
			".ts": api.LoaderTS,
			".js": api.LoaderJS,
		},
		EntryPoints: []string{path},
		Bundle:      true,
		Write:       false,
		GlobalName:  "k8x",
		Format:      api.FormatIIFE,
	}

	result := api.Build(options)

	for _, message := range result.Errors {
		log.Error().Msg(message.Text)
	}

	for _, message := range result.Warnings {
		log.Warn().Msg(message.Text)
	}

	if len(result.Errors) > 0 {
		log.Error().Msg("Cannot transform js")
	}

	if debug {
		log.Debug().Msg(string(result.OutputFiles[0].Contents))
	}

	return string(result.OutputFiles[0].Contents)
}

func injectEnv(vm *goja.Runtime) {
	m := make(map[string]interface{})

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if strings.Index(pair[0], "K8X_") == -1 {
			continue
		}

		m[pair[0]] = pair[1]
	}

	obj := vm.NewObject()

	err := obj.Set("$env", m)
	if err != nil {
		return
	}

	err = vm.Set("k8x", obj)

	if err != nil {
		log.Error().Msg("Cannot set $env variables")
		os.Exit(-1)
	}
}

// Executes tsx and returns its result
func Run(code string) any {
	vm := goja.New()

	injectEnv(vm)

	// Execute chart.tsx
	_, err := vm.RunString(code)

	if err != nil {
		log.Error().Msg("Can't evaluate chart:")
		log.Error().Err(err)
		os.Exit(-1)
	}

	k8x, ok := goja.AssertFunction(vm.Get("k8x").ToObject(vm).Get("default"))

	if !ok {
		log.Error().Err(err)
		os.Exit(-1)
	}

	obj, err := k8x(goja.Undefined())

	if err != nil {
		log.Error().Err(err)
		os.Exit(-1)
	}

	k8sExport, ok := obj.Export().(map[string]interface{})

	if !ok {
		log.Error().Msg("Cant cast to object")
		os.Exit(-1)
	}

	yml, _ := yaml.Marshal(k8sExport)
	fmt.Println(string(yml))

	return k8sExport
}
