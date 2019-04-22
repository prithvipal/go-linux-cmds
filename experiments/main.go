package main

import (
	"fmt"
	"io/ioutil"
	"syscall"
)

func main() {
	entries, err := ioutil.ReadDir("../")
	if err != nil {
		fmt.Println("Error", err)
	}
	for _, entry := range entries {
		stat := entry.Sys().(*syscall.Stat_t)
		fmt.Println(stat.Dev)
	}
}
