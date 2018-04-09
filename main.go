package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
)

func main() {
	fmt.Println("hello world")
}

func consulFunction() {
	config := consulapi.DefaultConfig()
	consul, _ := consulapi.NewClient(config)
	fmt.Println(consul)
}
