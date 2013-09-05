package hooks

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	WebHooksFilename = "playbill.webhooks"
	WebHookTimeout   = 2 * time.Second
)

// Custom client to enable canceling the request via the transport
var transport = http.Transport{}
var client = http.Client{
	Transport: &transport,
}

// Reads the web hooks file and returns a map of web hooks
func parseWebHooks() (handlers map[string][]string) {
	return parseHooksFile(WebHooksFilename)
}

// Send a POST request to the URL
func postWebHook(url string, data io.Reader, timeout time.Duration) {
	req, _ := http.NewRequest("POST", url, data)

	timer := time.AfterFunc(timeout, func() {
		transport.CancelRequest(req)
		log.Println(url, "timed out")
	})
	defer timer.Stop()

	client.Do(req)
}

// Trigger web hooks
func triggerWebHooks(name string, data io.Reader) int {
	handlers := parseWebHooks()
	urls, ok := handlers[name]

	// No urls available
	if !ok {
		return 0
	}

	n := len(handlers)

	wg := sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(n)

	for _, url := range urls {
		go func(url string) {
			postWebHook(url, data, WebHookTimeout)
			wg.Done()
		}(url)
	}

	return n
}
