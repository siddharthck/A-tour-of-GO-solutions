package main

import (
	"fmt"
	"sync"
)
// we need map with mutex for better concurrent use
type map_with_mutex struct {
	mutex sync.Mutex
	m map[string]int

}
//global var to store count the number of times url is visited
var url_map = map_with_mutex{m: make(map[string]int)}



type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.

// function to check if url is already visited
func Check_if_visited(url_map map_with_mutex , url string) bool{
	url_map.mutex.Lock()
	if (url_map.m[url] == 0){
		url_map.m[url]++
		url_map.mutex.Unlock()
		return false
	}
	url_map.mutex.Unlock()
	return true
}

func Crawl(url string, depth int, fetcher Fetcher) {

	
	 
	
	if depth <= 0 {
		
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		if !Check_if_visited(url_map,url){//if url not visited, print
		fmt.Println(err)
		}
		
		return
	}
	if !Check_if_visited(url_map,url){// if url not visited, print

	fmt.Printf("found: %s %q\n", url, body)

	}

	// channel for concurrent use (actually to wait untill each goroutine ends)
	ch:= make(chan int)

	for _, u := range urls {
		// you have to pass param u to give it to Crawl(inside closure)
		go func(u string){ 
			
		Crawl(u, depth-1, fetcher)
		ch <- 1
		}(u)
	
	}
	// you need to empty channel as many times you have used it to send
	for  range urls {
		
		<-ch
	}
	return
}



func main() {
	
	Crawl("https://golang.org/", 4, fetcher)

	
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
