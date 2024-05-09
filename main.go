package main

import (
	"fmt"
	"os"
	"path"
)

const colorCyanBold = "\033[1;36m"
const colorNone = "\033[0m"

var dirCount int32
var fileCount int32

// only path could work
func tree(basePath string, indent int) {
	f, err := os.Open(basePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fInfo, err := f.Readdir(-1)
	if err != nil {
		fmt.Println("Error occurred while reading directory:", err)
		os.Exit(1)
	}

	for _, file := range fInfo {
		fileName := file.Name()

		if file.IsDir() {
			dirCount += 1
			fmt.Printf("%s%*s\n", colorCyanBold, indent+len(fileName), fileName)
			tree(path.Join(basePath, fileName), indent+4)
		} else {
			fileCount += 1
			fmt.Printf("%s%*s\n", colorNone, indent+len(fileName), fileName)
		}
	}

	return
}

func main() {
	basePath := "/Users/bugraalparlsan/Desktop/tree-example"

	fmt.Printf("%s%s\n", colorCyanBold, basePath)
	tree(basePath, 4)

	fmt.Printf("\n%d directories, %d files\n", dirCount, fileCount)
}
