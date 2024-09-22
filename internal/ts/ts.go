package ts

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/evanw/esbuild/pkg/api"
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
		fmt.Println(message)
	}

	for _, message := range result.Warnings {
		fmt.Println(message)
	}

	if len(result.Errors) > 0 {
		fmt.Println("Cannot transform js")
	}

	if debug {
		fmt.Print(string(result.OutputFiles[0].Contents))
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

		m[strings.Replace(pair[0], "K8X_", "", 1)] = pair[1]
	}

	obj := vm.NewObject()

	err := obj.Set("$env", m)
	if err != nil {
		return
	}

	err = vm.Set("k8x", obj)

	if err != nil {
		panic(err)
	}
}

// Executes tsx and returns its result
func Run(code string) map[string]interface{} {
	vm := goja.New()

	injectEnv(vm)

	_, err := vm.RunString(code)

	if err != nil {
		fmt.Println("Can't evaluate chart:")
		os.Exit(-1)
	}

	k8x, ok := goja.AssertFunction(vm.Get("k8x").ToObject(vm).Get("default"))

	if !ok {
		panic("Please make sure you are exporting a function: export default () => ({})")
	}

	obj, err := k8x(goja.Undefined())

	if err != nil {
		panic(err)
	}

	k8sExport, ok := obj.Export().(map[string]interface{})

	if !ok {
		panic("Cant cast to object")
	}

	ns, ok := k8sExport["namespace"].(string)

	// Patching namespaces
	if ns != "" && ok {
		for _, component := range k8sExport["components"].([]interface{}) {
			if component == nil {
				continue
			}

			comp := component.(map[string]interface{})
			metadata := comp["metadata"].(map[string]interface{})
			metadata["namespace"] = ns
		}
	}

	return k8sExport
}
