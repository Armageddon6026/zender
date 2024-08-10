package main

import (
	"log"

	"github.com/Armageddon6026/zender/pkg/common"
	"github.com/Armageddon6026/zender/pkg/server"
)

func main() {
	config, err := common.GetSystemConfig()
	if err != nil {
		log.Fatal(err)
	}

	server := server.New(config)
	server.Run()
}
