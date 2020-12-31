package main

import (
	"apascualco.com/api/server"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() { server.Start() }()
	wg.Wait()
}
