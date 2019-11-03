package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Hello, TechBash!")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go DoImportantWork(i, rand.Intn(5), &wg)
	}

	fmt.Println("Ending the application")
	wg.Wait()
}

func DoImportantWork(id int, sleepTime int, wg *sync.WaitGroup) {
	fmt.Printf("%d > Starting to do work (Sleeptime=%d\n", id, sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Printf("%d > Done doing work\n", id)
	wg.Done()
}
