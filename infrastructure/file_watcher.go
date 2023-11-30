package infrastructure

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/lucaschain/beholder/core"
)

func loop(watcher *fsnotify.Watcher, callback core.ChangeCallback) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				continue
			}
			changeEvent := core.ChangeEvent{Type: event.Op.String(), FileName: event.Name}
			callback(&changeEvent, nil)
		case err, _ := <-watcher.Errors:
			if err != nil {
				callback(nil, &err)
			}
		}
	}
}

func FileWatcher(paths []string, callback core.ChangeCallback) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		callback(nil, &err)
	}
	defer watcher.Close()

	go func() {
		loop(watcher, callback)
	}()

	for _, path := range paths {
		err = watcher.Add(path)
		if err != nil {
			log.Fatal(err)
		}
	}

	<-make(chan struct{})
}
