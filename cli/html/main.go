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
	str, e := engine.RunHTMLFile(engine.NewScope(engine.NewGlobalScope(dir, engine.DefaultReadFileFunc)), file)

	if e != nil {
		fmt.Println("Error!")
		panic(e.Error())
	}

	fmt.Println(str)
}
