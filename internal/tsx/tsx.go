package tsx

import (
	"encoding/json"
	"fmt"
	"github.com/dop251/goja"
	"github.com/evanw/esbuild/pkg/api"
	"os"
	"strings"
)

type Object struct {
	Id       string
	Props    map[string]interface{}
	Children []interface{}
}

// Loads and transpiles tsx files
func Load(path string) string {

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
		fmt.Println(message)
	}

	for _, message := range result.Warnings {
		fmt.Println(message)
	}

	if len(result.Errors) > 0 {
		panic("")
	}

	return string(result.OutputFiles[0].Contents)
}

// Executes tsx and returns its result
func Run(code string) string {
	vm := goja.New()

	fc := func(id string, props map[string]interface{}, children ...interface{}) Object {

		if strings.Contains(id, "__jsx(") {
			obj, err := vm.RunString(id)

			if err != nil {
				panic(err)
			}

			sum, ok := goja.AssertFunction(obj)

			if !ok {
				panic("Not a function")
			}

			val, err := sum(goja.Undefined(), vm.ToValue(props), vm.ToValue(children))

			if err != nil {
				panic(err)
			}

			return val.Export().(Object)
		}

		return Object{id, props, children}
	}
	err := vm.Set("__jsx", fc)

	if err != nil {
		panic(err)
	}

	m := make(map[string]interface{})

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if strings.Index(pair[0], "K8X_") == -1 {
			continue
		}

		m[pair[0]] = pair[1]
	}

	err = vm.Set("$env", m)

	if err != nil {
		panic(err)
	}
	_, err = vm.RunString(code)

	if err != nil {
		panic(err)
	}

	sum, ok := goja.AssertFunction(vm.Get("k8x").ToObject(vm).Get("default"))

	if !ok {
		panic("Not a function")
	}

	obj, err := sum(goja.Undefined())

	if err != nil {
		panic(err)
	}

	num, ok := obj.Export().(Object)

	if !ok {
		panic("Cant cast result to Object")
	}

	yml, err := json.Marshal(num)

	if err != nil {
		panic(err)
	}

	return string(yml)
}
