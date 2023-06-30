package main

import (
	"fmt"
	"os"
	"path"

	"github.com/MarionetteEnsemble/Puppeto/engine"
)

func main() {
	file := os.Args[1]
	dir := path.Dir(file)
	v, e := engine.RunFile(engine.NewScope(engine.NewGlobalScope(dir, engine.DefaultReadFileFunc)), file)

	if e != nil {
		fmt.Println("Error!")
		fmt.Println(engine.StringifyError(engine.Stack{}, e))
	} else {
		fmt.Println("Last Value: " + engine.ToStr(v))
	}
}
