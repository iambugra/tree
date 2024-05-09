package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	argCount := len(os.Args)
	fmt.Printf("Number of arguments: %d\n", argCount)

	hiddenFlag := flag.Bool("a", false, "All files are printed.")
	onlyDirFlag := flag.Bool("d", false, "List directories only.")
	maxDepthFlag := flag.Int("max-depth", 8, "Maximum depth to list files.")
	flag.Parse()

	fmt.Println("a: ", *hiddenFlag)
	fmt.Println("d: ", *onlyDirFlag)
	fmt.Println("max-depth: ", *maxDepthFlag)
	fmt.Println("tails: ", flag.Args())
}
