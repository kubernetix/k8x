package tsx

import (
	"github.com/dop251/goja"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

type Object struct {
	Id       string
	Props    map[string]interface{}
	Children []interface{}
}

// Loads and transpiles tsx files
func Load(path string, debug bool) string {

	options := api.BuildOptions{
		Loader: map[string]api.Loader{
			".tsx": api.LoaderTSX,
		},
		EntryPoints: []string{path},
		JSX:         api.JSXTransform,
		Bundle:      true,
		Write:       false,
		JSXFragment: "__jsxFragment",
		JSXFactory:  "__jsx",
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

	err := vm.Set("$env", m)

	if err != nil {
		log.Error().Msg("Cannot set $env variables")
		os.Exit(-1)
	}
}

func injectJsxFunctions(vm *goja.Runtime) {
	fc := func(id string, props map[string]interface{}, children ...interface{}) Object {

		if strings.Contains(id, "__jsx(") {
			obj, err := vm.RunString(id)

			if err != nil {
				panic(err)
			}

			sum, ok := goja.AssertFunction(obj)

			if !ok {
				log.Error().Msg("Provided jsx callback is not a function")
				os.Exit(-1)
			}

			val, err := sum(goja.Undefined(), vm.ToValue(props), vm.ToValue(children))

			if err != nil {
				log.Error().Err(err)
				os.Exit(-1)
			}

			return val.Export().(Object)
		}

		return Object{id, props, children}
	}
	err := vm.Set("__jsx", fc)

	if err != nil {
		log.Error().Err(err)
		os.Exit(-1)
	}
}

// Executes tsx and returns its result
func Run(code string) Object {
	vm := goja.New()

	injectJsxFunctions(vm)
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

	num, ok := obj.Export().(Object)

	if !ok {
		log.Error().Msg("Cant cast to object")
		os.Exit(-1)
	}

	return num
}
