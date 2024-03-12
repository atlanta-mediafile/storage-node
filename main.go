package main

import (
	"fmt"
	"mediafile/storage_node/config"
)

func main() {
	node, err := config.GetConfig()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v", node)
	}

}
