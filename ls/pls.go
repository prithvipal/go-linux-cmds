package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"syscall"
)

type lsStat struct {
	mode         os.FileMode
	nLink        uint16
	username     string
	groupname    string
	size         int64
	lastModified string
	name         string
}

func (s *lsStat) print() {
	//fmt.Println("What is this?")
	fmt.Printf("%v %2v %v %v %4v %v %v\n", s.mode, s.nLink, s.username, s.groupname, s.size, s.lastModified, s.name)
}

func getStats(entries []os.FileInfo) []lsStat {
	stats := make([]lsStat, 0)
	for _, entry := range entries {
		ls := lsStat{mode: entry.Mode(),
			size:         entry.Size(),
			lastModified: entry.ModTime().Format("Apr _2 15:04"),
			name:         entry.Name()}
		stat := entry.Sys().(*syscall.Stat_t)
		ls.nLink = stat.Nlink
		newUser, err := user.LookupId(fmt.Sprint(stat.Uid))
		if err == nil {
			ls.username = newUser.Username
		}

		group, err := user.LookupGroupId(fmt.Sprint(stat.Gid))
		if err == nil {
			ls.groupname = group.Name
		}
		stats = append(stats, ls)

	}
	return stats
}

func printStats(lsStats []lsStat) {
	for _, ls := range lsStats {
		ls.print()
	}
}

func getFilePath() string {
	var filepath string
	if len(os.Args) == 1 {
		filepath = "./"
	} else if len(os.Args) == 2 {
		filePath := os.Args[1]
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Println("Error", err)
			fmt.Printf("ls: %v: No such file or directory\n", filePath)
			os.Exit(1)
		}
		return filePath
	} else {
		fmt.Println("Invalid command argument. Valid example:")
		fmt.Println("./ls <file or dir path>")
		os.Exit(1)
	}
	return filepath
}

func getEntries() []os.FileInfo {
	filepath := getFilePath()
	var entries []os.FileInfo
	entries, err := ioutil.ReadDir(filepath)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	return entries
}

func main() {
	entries := getEntries()
	lsStats := getStats(entries)
	printStats(lsStats)

}
