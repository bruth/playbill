package hooks

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
)

const LocalHooksFilename = "playbill.hooks"

func parseLocalHooks() map[string][]string {
	return parseHooksFile(LocalHooksFilename)
}

func execCmdHook(filename string, data io.Reader) {
	cmd := exec.Command(filename)
	cmd.Stdin = data
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

// Trigger local hooks
func triggerLocalHooks(name string, data io.Reader) int {
	handlers := parseWebHooks()
	filenames, ok := handlers[name]

	// No urls available
	if !ok {
		return 0
	}

	n := len(filenames)

	wg := sync.WaitGroup{}
	wg.Add(n)
	defer wg.Wait()

	// TODO add error handling?
	for _, filename := range filenames {
		go func(filename string) {
			execCmdHook(filename, data)
			wg.Done()
		}(filename)
	}

	return n
}

// NewLocalHooks creates a playbill.hooks file if one does not already exist
func NewLocalHooks() error {
	f, err := OpenLocalHooks()
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	return ioutil.WriteFile(f.Name(), nil, 0666)
}

func OpenLocalHooks() (*os.File, error) {
	err := NewLocalHooks()
	if err != nil {
		return nil, err
	}
	cwd, _ := os.Getwd()
	fn := path.Join(cwd, LocalHooksFilename)
	return os.Open(fn)
}

func AppendLocalHooks(pairs [][]string) error {
	f, err := OpenLocalHooks()
	if err != nil {
		return err
	}
	defer f.Close()
	for _, p := range pairs {
		line := strings.Join(p, "\t")
		f.WriteString(fmt.Sprintf("%s\n", line))
	}
	return err
}
