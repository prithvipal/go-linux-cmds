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
		fmt.Println("----------------")
	}
}
