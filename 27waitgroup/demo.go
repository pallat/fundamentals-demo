package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sync"
)

var needed int
var mux = &sync.Mutex{}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(100)
	// go works("go", wg)
	// go works("Rust", wg)
	// go works("React", wg)

	for i := 0; i < 100; i++ {
		go racing(i, wg)
	}

	wg.Wait()

}

func racing(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	mux.Lock()
	defer mux.Unlock()
	if i%2 == 0 {
		needed = i
	} else {
		fmt.Println(needed)
	}
}

func works(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	b := sha256.Sum256([]byte(s))
	fmt.Println(
		base64.StdEncoding.EncodeToString(b[:]),
	)
}
