package main

import (
	"fmt"
	"time"
	"sync"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main(){
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTempo total de execução: %v", time.Since(t0))
	fmt.Printf("\nOs resultados são %v", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("O resultado da base dados é ", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}