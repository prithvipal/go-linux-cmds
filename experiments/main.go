package main

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"syscall"
)

func main() {
	entries, err := ioutil.ReadDir("../")
	if err != nil {
		fmt.Println("Error", err)
	}
	for _, entry := range entries {
		//Getting permissions
		fmt.Println(entry.Mode())
		stat := entry.Sys().(*syscall.Stat_t)
		// Getting number of links
		fmt.Println(stat.Nlink)
		//Getting Username
		newUser, err := user.LookupId(fmt.Sprint(stat.Uid))
		if err == nil {
			fmt.Println(newUser.Username)
		}

		//Getting Group name
		group, err := user.LookupGroupId(fmt.Sprint(stat.Gid))
		if err == nil {
			fmt.Println(group.Name)
		}

		//Getting filesize
		fmt.Println(entry.Size())

		//Getting last modified
		//t := time.Now()
		//fmt.Println(t.Format("Apr _2 15:04"))
		fmt.Println(entry.ModTime().Format("Apr _2 15:04"))

		//Getting file name
		fmt.Println(entry.Name())

		fmt.Println("----------------")

	}
}
