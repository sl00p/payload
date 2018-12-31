package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func checkAlive(conf string, host string, maxThread int64, sleep int64, statusCode int) {
	var (
		wg        sync.WaitGroup
		wRoutines = make(chan struct{}, maxThread)
	)
	f, err := os.Open(conf)
	if err != nil {
		fmt.Println("[E]: Can't open file", conf)
		return
	}
	defer f.Close()

	br := bufio.NewReader(f)
	for {
		s, _, r := br.ReadLine()
		if r == io.EOF {
			break
		}
		wRoutines <- struct{}{}
		wg.Add(1)
		go func(s string, host string, sleep int64, statusCode int) {
			var url = host + "/" + s
			time.Sleep(time.Duration(sleep) * time.Second)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("[E]:", err.Error())
				<-wRoutines
				wg.Done()
				return
			}

			defer resp.Body.Close()
			if resp.StatusCode == statusCode {
				fmt.Println("[I]:", url, "is hit.")
			}
			<-wRoutines
			wg.Done()
		}(string(s), host, sleep, statusCode)
	}
	wg.Wait()
}

func main() {
	var (
		config     string
		maxThread  int64
		host       string
		statusCode int
		sleep      int64
	)
	flag.StringVar(&config, "c", "conf", "Default config path.")
	flag.Int64Var(&maxThread, "n", 100, "Multi-thread number.")
	flag.StringVar(&host, "u", "", "Host url")
	flag.Int64Var(&sleep, "t", 0, "Sleep time, default 0.")
	flag.IntVar(&statusCode, "s", 200, "Status code, default 200.")
	flag.Parse()
	if len(host) <= 0 {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		return
	}
	checkAlive(config, host, maxThread, sleep, statusCode)
}
