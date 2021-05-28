package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			df, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i))
			if err != nil {
				log.Fatal(err)
			}
			byte, err := io.ReadAll(df.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(byte))
			wg.Done()
		}(i)
	}
	wg.Wait()

}
