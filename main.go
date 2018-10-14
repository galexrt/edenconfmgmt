package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/hashicorp/memberlist"
)

func main() {
	count, _ := strconv.Atoi(os.Args[1])

	config := memberlist.DefaultLocalConfig()
	config.BindAddr = fmt.Sprintf("127.0.0.%d", count)
	config.AdvertiseAddr = fmt.Sprintf("127.0.0.%d", count)
	config.Name = fmt.Sprintf("node%d", count)
	list, err := memberlist.Create(config)
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	members := []string{}

	for i := 1; i < count; i++ {
		if i != count {
			members = append(members, fmt.Sprintf("127.0.0.%d", i))
		}
	}

	fmt.Printf("Trying to join members: %+v\n", members)

	// Join an existing cluster by specifying at least one known member.
	_, err = list.Join(members)
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}

	if len(members) == 1 {
		list.PushPull()
	}

	for {
		fmt.Printf("--- %d ---\n", list.NumMembers())
		// Ask for members of the cluster
		for _, member := range list.Members() {
			fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
		}
		<-time.After(3 * time.Second)
	}
}
