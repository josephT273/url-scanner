package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(url)

	if err != nil {
		fmt.Printf("❌ %s (error: %v)\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("✅ %s (status: %d)\n", url, resp.StatusCode)
	} else {
		fmt.Printf("❌ %s (status: %d)\n", url, resp.StatusCode)

	}

}

func main() {
	urls := flag.String("w", "urls.txt", "Setup wordlist to scan the websites is live or not")
	flag.Parse()

	data, err := os.Open(*urls)
	check(err)
	defer data.Close()

	var wg sync.WaitGroup
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		url := scanner.Text()
		wg.Add(1)
		go checkURL(url, &wg)
	}
	wg.Wait()

}
