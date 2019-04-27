package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func validate() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid command arguments")
		fmt.Println("Ex: ./cp <src filepath> <des file or dir path>")
		os.Exit(1)
	}
}

func getSrcPath() string {
	path := os.Args[1]
	file, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if !file.Mode().IsRegular() {
		fmt.Printf("%v is not a regular file\n", path)
		os.Exit(1)
	}
	return path
}

func getDestPath() string {
	path := os.Args[2]
	dstfile, err := os.Stat(path)
	if os.IsNotExist(err) {
		dstdirpath := filepath.Dir(path)
		if dstdir, err := os.Stat(dstdirpath); os.IsNotExist(err) || !dstdir.Mode().IsDir() {
			fmt.Printf("Invalide destination path %v", path)
			os.Exit(1)
		}
		return path
	} else if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	} else if dstfile.IsDir() {
		srcfile, _ := os.Stat(os.Args[1])
		dstfilepath := filepath.Join(path, srcfile.Name())
		return dstfilepath
	}
	return path
}

func getPaths() (string, string) {
	srcpath := getSrcPath()
	dstpath := getDestPath()
	return srcpath, dstpath
}

func main() {
	validate()
	srcpath, dstpath := getPaths()
	fmt.Println(srcpath)
	fmt.Println(dstpath)
	//wd, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	//fmt.Println(wd)
	//parent := filepath.Dir("/Users/prithvipalsingh/PRITHVI/Github/src/go-linux-cmds11/README.msds")
	//fmt.Println(parent)
}
