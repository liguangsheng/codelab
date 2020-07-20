package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	cnt, err := strconv.ParseFloat(os.Args[1], 32)
	if err != nil {
		panic(err)
	}
	ts := time.Duration(1000000.0/cnt) * time.Microsecond
	for i := 0; i < int(cnt); i++ {
		go func() {
			resp, err := http.Get("http://:8080/ping")
			if err != nil {
				fmt.Println(err)
			}
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
		}()
		time.Sleep(ts)
	}
}
