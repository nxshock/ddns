package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	flagWorkingDir  string
	flagHTTPTimeout time.Duration

	wg = new(sync.WaitGroup)

	httpClient *http.Client

	errLog = log.New(os.Stderr, "", 0)
	msgLog = log.New(os.Stdout, "", 0)
)

func main() {
	files, err := ioutil.ReadDir(flagWorkingDir)
	if err != nil {
		errLog.Fatalln(err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".toml") {
			updater, err := LoadUpdater(filepath.Join(flagWorkingDir, file.Name()))
			if err != nil {
				errLog.Printf("error while reading %s: %v", file.Name(), err)
				continue
			}
			wg.Add(1)
			go executeUpdate(updater)
		}
	}
	wg.Wait()
}

func executeUpdate(updater *Updater) {
	defer wg.Done()

	result, err := updater.Update()
	if err != nil {
		errLog.Printf("error while updating %s: %v", updater.Name, err)
		return
	}

	msgLog.Printf("%s: %v", updater.Name, result)
}
