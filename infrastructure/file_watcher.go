package infrastructure

import (
	"context"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/lucaschain/beholder/core"
	"github.com/lucaschain/beholder/core/event_types"
)

func loop(watcher *fsnotify.Watcher, callback core.ChangeCallback) *error {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				continue
			}
			changeEvent := core.ChangeEvent{
				Type:     event_types.FromString(event.Op.String()),
				FileName: event.Name,
			}
			err := callback(&changeEvent, nil)
			if err != nil {
				return err
			}
		case err, _ := <-watcher.Errors:
			if err != nil {
				callback(nil, &err)
			}
		}
	}
}

func FileWatcher(paths []string, callback core.ChangeCallback, ctx context.Context) *error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		callback(nil, &err)
	}
	defer watcher.Close()

	go func() {
		commandError := loop(watcher, callback)

		if commandError != nil {
			ctx.Done()
		}
	}()

	for _, path := range paths {
		err = watcher.Add(path)
		if err != nil {
			log.Fatal(err)
		}
	}

	<-ctx.Done()

	err = ctx.Err()
	return &err
}
