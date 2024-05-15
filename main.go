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

var boxLines BoxDrawings

const colorCyanBold = "\033[1;36m"
const colorNone = "\033[0m"

var dirCount int32
var fileCount int32
var hiddenFlag *bool

func tree(basePath string, indent string) {
	f, err := os.Open(basePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	printLine(fileInfo, indent)
	if !fileInfo.IsDir() {
		return
	}

	fileInfos, err := f.Readdir(-1)
	if err != nil {
		fmt.Println("Error occurred while reading directory:", err)
		os.Exit(1)
	}

	lenFiles := len(fileInfos)
	twoHorizontalLines := boxLines.horizontal + boxLines.horizontal
	var connectingLine string
	var modifiedIndent string

	for i, file := range fileInfos {
		fileName := file.Name()

		if i == lenFiles-1 {
			connectingLine = boxLines.upAndRight
		} else {
			connectingLine = boxLines.verticalAndRight
		}

		lenIndent := len(indent) / 3 // Every kind of line in boxLines is 3 bytes long.
		//fmt.Printf("%d len horizontal line \n", len(boxLines.horizontal))
		//fmt.Printf("%d len indent\n", lenIndent)
		//fmt.Printf("Provided indent is |%s| for file %s\n", indent, fileName)
		if lenIndent < 3 {
			//fmt.Printf("len short if\n")
			modifiedIndent = indent
		} else if string([]rune(indent)[lenIndent-3]) == boxLines.upAndRight {
			//fmt.Printf("Upright if\n")
			modifiedIndent = indent[:lenIndent-3] + "   "
		} else if string([]rune(indent)[lenIndent-3]) == boxLines.verticalAndRight {
			//fmt.Printf("verticalAndRight if\n")
			modifiedIndent = indent[:lenIndent-3] + boxLines.vertical + "  "
		}

		tree(path.Join(basePath, fileName), modifiedIndent+connectingLine+twoHorizontalLines)
	}

	return
}

func printLine(fileInfo os.FileInfo, indent string) {
	var color string
	if fileInfo.IsDir() {
		dirCount += 1
		color = colorCyanBold
	} else {
		fileCount += 1
		color = colorNone
	}

	if *hiddenFlag {
		fmt.Printf("%s%s%s\n", color, indent, fileInfo.Name())
	} else if !strings.HasPrefix(fileInfo.Name(), ".") {
		fmt.Printf("%s%s%s\n", color, indent, fileInfo.Name())
	}
}

func formattedPrint(fileName string, indent string, color string) {
	if *hiddenFlag {
		fmt.Printf("%s%s%s\n", color, indent, fileName)
	} else if !strings.HasPrefix(fileName, ".") {
		fmt.Printf("%s%s%s\n", color, indent, fileName)
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

	// set BoxDrawings struct
	setBoxLines(&boxLines)

	if len(flag.Args()) == 0 {
		fmt.Println("No arguments given. Give exactly one directory.")
		os.Exit(1)
	}
	if len(flag.Args()) > 1 {
		fmt.Println("Too many arguments. Give exactly one directory.")
		os.Exit(1)
	}
	basePath := flag.Args()[0]

	tree(basePath, "")

	fmt.Printf("\n%d directories, %d files\n", dirCount, fileCount)
}
