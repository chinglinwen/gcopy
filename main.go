package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var example = `
Examples: 
	gcopy src dst
	gcopy src1 src2 dst

`

func main() {
	flag.Usage = func() {
		fmt.Println("Usage of gcopy:")
		flag.PrintDefaults()
		fmt.Printf(example)
	}
	verbose := flag.Bool("v", false, "verbose")
	version := flag.Bool("version", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println("version=1.0.0, 2017-1-19")
		os.Exit(1)
	}

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("no src or dst specified")
		os.Exit(1)
	}

	last := len(args) - 1
	dst := args[last]

	for i := 0; i < last; i++ {
		src := args[i]
		if *verbose {
			fmt.Printf("copying... %v to %v\n", src, dst)
		}
		checkErr(docp(src, dst))
	}
}

func docp(src, dst string) error {
	file, err := os.Open(src)
	checkErr(err)
	m, err := file.Stat()
	checkErr(err)
	filetype := m.Mode()
	if filetype.IsRegular() {
		return CopyFile(src, dst)
	}
	if filetype.IsDir() {
		return CopyDir(src, dst)
	}
	return fmt.Errorf("%v is unexpected type\n", src)
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
