package core

import (
	"github.com/fsnotify/fsnotify"
)

type Callback func(*fsnotify.Event, *error)

func loop(watcher *fsnotify.Watcher, callback Callback) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				continue
			}
			callback(&event, nil)
		case err, _ := <-watcher.Errors:
			if err != nil {
				callback(nil, &err)
			}
		}
	}
}

func Start(path string, callback Callback) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		callback(nil, &err)
	}
	defer watcher.Close()

	go func() {
		loop(watcher, callback)
	}()

	err = watcher.Add(path)
	if err != nil {
		callback(nil, &err)
	}

	<-make(chan struct{})
}
