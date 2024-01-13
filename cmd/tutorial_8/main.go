package main

import (
	"fmt"
	"sync"
	"time"
)

var aMutex = sync.RWMutex{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}
var aWaitGroup = sync.WaitGroup{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		aWaitGroup.Add(1)
		go dbCall(i)
	}
	aWaitGroup.Wait()
	fmt.Printf("\nTotal executrion time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(index int) {
	var delay float32 = 2000
	// var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[index])
	save(dbData[index])
	log()
	// aMutex.Lock()
	// results = append(results, dbData[index])
	// aMutex.Unlock()
	aWaitGroup.Done()
}

func save(result string) {
	aMutex.Lock()
	results = append(results, result)
	aMutex.Unlock()
}

func log() {
	aMutex.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	aMutex.RUnlock()
}
