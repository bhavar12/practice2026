package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	urls := []string{"abc", "pqr"}
	fmt.Println(getUrls(urls))
	time.Sleep(1 * time.Second)
}

func getUrls(urls []string) []error {
	errors := make([]error, len(urls))
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(urls); i++ {
			errors[i] = fmt.Errorf("url %s", urls[i])
		}
	}()
	wg.Wait()
	return errors
}
