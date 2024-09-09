package main

import (
	"encoding/json"
	"fmt"
	"github.com/dop251/goja"
	"github.com/evanw/esbuild/pkg/api"
	"os"
	"runtime"
	"strings"
)

type Object struct {
	Id       string
	Props    map[string]interface{}
	Children []interface{}
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("ERROR: Cant find chart. Usage: k8x chart.tsx")
		os.Exit(-1)
	}

	runtime.LockOSThread()

	path := os.Args[1]

	result := api.Build(api.BuildOptions{
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
	})

	for _, message := range result.Errors {
		fmt.Println(message)
	}

	for _, message := range result.Warnings {
		fmt.Println(message)
	}

	if len(result.Errors) > 0 {
		panic("")
	}

	code := string(result.OutputFiles[0].Contents)

	gojajs(code)
	// quick(code)
}

func gojajs(code string) {
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

	fmt.Println(string(yml))
}
