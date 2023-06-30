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
		if o, ok := e.(engine.PObject); ok {
			v, e := o["toString"].(engine.PFunction)(engine.Stack{}, engine.PArray{})

			if e != nil {
				panic(e)
			}

			panic(v)
		} else {
			panic(e)
		}
	} else {
		fmt.Println("Last Value: " + engine.ToStr(v))
	}
}
