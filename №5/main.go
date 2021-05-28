package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {

	for i := 1; i <= 5; i++ {
		df, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		byte, err := io.ReadAll(df.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(byte))
		err = ioutil.WriteFile("storage/posts/"+strconv.Itoa(i)+"test.txt", byte, 0666)
		if err != nil {
			log.Fatal(err)
		}
	}
}
