package main

import (
	"log"
	"mediafile/storage_node/config"
	"mediafile/storage_node/server"
	"sync"
)

func main() {

	node, err := config.GetConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}

	//fmt.Printf("%#v\n", node)

	var wg sync.WaitGroup

	wg.Add(1)

	go server.StartGrpcServer(&node)
	defer wg.Done()

	wg.Wait()
}
