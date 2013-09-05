package hooks

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
)

// Triggers the handlers associated with the hook
func Trigger(name string, payload interface{}) (int, error) {
	raw := map[string]interface{}{
		"hook":    name,
		"payload": payload,
	}

	// Encode the data as JSON for all hook handlers to receive
	var data bytes.Buffer
	encoder := json.NewEncoder(&data)
	err := encoder.Encode(raw)

	if err != nil {
		return 0, err
	}

	var wc, lc int
	wg := sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(2)

	// Trigger hooks asynchronously
	go func() {
		wc = triggerWebHooks(name, &data)
		wg.Done()
	}()
	go func() {
		lc = triggerLocalHooks(name, &data)
		wg.Done()
	}()

	return wc + lc, nil
}

// Parse hooks files which contains two columns delimited by a tab. The first
// column is the hook name and the second is the target which is handled
// downstream
func parseHooksFile(filename string) (handlers map[string][]string) {
	cwd, _ := os.Getwd()
	p := path.Join(cwd, filename)
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return nil
	}

	// Create scanner
	buf := bytes.NewBuffer(b)
	scanner := bufio.NewScanner(buf)

	// Scan lines and split on tabs
	for scanner.Scan() {
		toks := strings.SplitN(scanner.Text(), "\t", 1)
		arr, ok := handlers[toks[0]]
		if !ok {
			arr = []string{}
		}
		handlers[toks[0]] = append(arr, toks[1])
	}
	return nil
}
