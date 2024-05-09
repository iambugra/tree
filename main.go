package main

import (
	"fmt"
	"os"
	"path"
)

// only path could work
func tree(basePath string, indent int) {
	f, err := os.Open(basePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fInfo, err := f.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, file := range fInfo {
		fileName := file.Name()
		fmt.Printf("%*s\n", indent+len(fileName), fileName)

		if file.IsDir() {
			tree(path.Join(basePath, fileName), indent+4)
		}
	}

	return

	//fileName := fileInfo.Name()
	//if !fileInfo.IsDir() {
	//	fmt.Printf("%*s", indent+len(fileName), fileName)
	//	return
	//} else {
	//	fmt.Printf("%*s", indent+len(fileName), fileName)
	//
	//}
}

func main() {
	basePath := "/Users/bugraalparlsan/Desktop/tree-example"

	tree(basePath, 0)
}
