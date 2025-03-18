package multithreading

import (
	"fmt"
	"net/http"
	"sync"
)

func FetchURL() {
	urls := []string{
		"https://golang.org",
		"https://go.dev/blog/pipelines",
	}

	channel := make(chan string, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := http.Get(url)
			if err != nil {
				channel <- fmt.Sprintf("Error fetching %s: %v", url, err)
				return
			}
			channel <- fmt.Sprintf("Fetched %s succesfully", url)
		}()
	}

	wg.Wait()
	close(channel)
	fmt.Println("All URLs have been fetched")

	for msg := range channel {
		fmt.Println(msg)
	}
}
