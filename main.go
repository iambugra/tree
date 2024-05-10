package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

type BoxDrawings struct {
	upAndRight       string
	vertical         string
	horizontal       string
	verticalAndRight string
}

const colorCyanBold = "\033[1;36m"
const colorNone = "\033[0m"

var dirCount int32
var fileCount int32
var hiddenFlag *bool

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
			formattedPrint(fileName, indent, colorCyanBold)
			tree(path.Join(basePath, fileName), indent+4)
		} else {
			fileCount += 1
			formattedPrint(fileName, indent, colorNone)
		}
	}

	return
}

func formattedPrint(fileName string, indent int, color string) {
	if *hiddenFlag {
		fmt.Printf("%s%*s\n", color, indent+len(fileName), fileName)
	} else if !strings.HasPrefix(fileName, ".") {
		fmt.Printf("%s%*s\n", color, indent+len(fileName), fileName)
	}
}

func setBoxLines(box *BoxDrawings) {
	box.upAndRight = fmt.Sprintf("%c", 9492)
	box.vertical = fmt.Sprintf("%c", 9474)
	box.horizontal = fmt.Sprintf("%c", 9472)
	box.verticalAndRight = fmt.Sprintf("%c", 9500)
}

func main() {
	hiddenFlag = flag.Bool("a", false, "All files are listed.")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("No arguments given. Give exactly one directory.")
		os.Exit(1)
	}
	if len(flag.Args()) > 1 {
		fmt.Println("Too many arguments. Give exactly one directory.")
		os.Exit(1)
	}
	basePath := flag.Args()[0]

	fmt.Printf("%s%s\n", colorCyanBold, basePath)
	tree(basePath, 4)

	fmt.Printf("\n%d directories, %d files\n", dirCount, fileCount)
}
