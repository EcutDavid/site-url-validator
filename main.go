// Simply read a bunch of URL from stdin or os args and validate their state.
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func fetch(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("got error %v", err)
		return
	}

	ch <- fmt.Sprintf("got status code %v", resp.StatusCode)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	urls := []string{}
	if len(os.Args) > 1 {
		urls = os.Args[1:]
	} else {
		fmt.Println("Reading URL list from stdin")
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
	}

	if len(urls) == 0 {
		fmt.Println("No URL received")
		os.Exit(0)
	}

	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}

	for i := range urls {
		fmt.Printf("%v: %v\n", urls[i], <-ch)
	}
}
