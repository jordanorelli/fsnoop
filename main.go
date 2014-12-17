package main

import (
	"code.google.com/p/go.exp/inotify"
	"flag"
	"fmt"
	"os"
)

var options struct {
	track trackMask
}

func main() {
	flag.Parse()
	if options.track == 0 {
		options.track = trackMask(inotify.IN_ALL_EVENTS)
	}

	w, err := inotify.NewWatcher()
	if err != nil {
		bail(1, "couldn't create new watcher: %v", err)
	}

	if flag.NArg() > 0 {
		for _, path := range flag.Args() {
			if err := w.AddWatch(path, uint32(options.track)); err != nil {
				// if err := w.Watch(d); err != nil {
				bail(1, "couldn't add watch: %v", err)
			}
		}
	} else {
		d, err := os.Getwd()
		if err != nil {
			bail(1, "couldn't get current directory: %v", err)
		}
		if err := w.AddWatch(d, uint32(options.track)); err != nil {
			// if err := w.Watch(d); err != nil {
			bail(1, "couldn't add watch: %v", err)
		}
	}

	cookies := make(map[uint32]*inotify.Event, 4)

	for {
		select {
		case ev := <-w.Event:
			if ev.Mask&inotify.IN_ACCESS > 0 {
				fmt.Printf("access %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_ATTRIB > 0 {
				fmt.Printf("attrib %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_CLOSE > 0 {
				fmt.Printf("close %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_CLOSE_NOWRITE > 0 {
				fmt.Printf("close-nowrite %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_CLOSE_WRITE > 0 {
				fmt.Printf("close-write %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_CREATE > 0 {
				fmt.Printf("create %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_DELETE > 0 {
				fmt.Printf("delete %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_DELETE_SELF > 0 {
				fmt.Printf("delete-self %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_DONT_FOLLOW > 0 {
				fmt.Printf("dont-follow %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_IGNORED > 0 {
				fmt.Printf("ignored %s\n", ev.Name)
			}
			// if ev.Mask & inotify.IN_ISDIR > 0 {
			//     fmt.Printf("isdir %s\n", ev.Name)
			// }
			if ev.Mask&inotify.IN_MODIFY > 0 {
				fmt.Printf("modify %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_MOVE > 0 {
				prev, ok := cookies[ev.Cookie]
				if !ok {
					cookies[ev.Cookie] = ev
				} else {
					fmt.Printf("move %s -> %s\n", prev.Name, ev.Name)
					delete(cookies, ev.Cookie)
				}
			}
			if ev.Mask&inotify.IN_MOVED_FROM > 0 {
				fmt.Printf("moved-from %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_MOVED_TO > 0 {
				fmt.Printf("moved-to %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_MOVE_SELF > 0 {
				fmt.Printf("move-self %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_ONESHOT > 0 {
				fmt.Printf("oneshot %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_ONLYDIR > 0 {
				fmt.Printf("onlydir %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_OPEN > 0 {
				fmt.Printf("open %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_Q_OVERFLOW > 0 {
				fmt.Printf("q-overflow %s\n", ev.Name)
			}
			if ev.Mask&inotify.IN_UNMOUNT > 0 {
				fmt.Printf("unmount %s\n", ev.Name)
			}
		case err := <-w.Error:
			bail(1, "error on watcher: %v", err)
		}
	}
}

func init() {
	flag.Var(&options.track, "track", "comma-separated list of events to track")
}
