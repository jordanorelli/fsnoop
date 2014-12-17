package main

import (
	"code.google.com/p/go.exp/inotify"
	"fmt"
	"strings"
)

type trackMask uint32

func (t *trackMask) String() string {
	parts := make([]string, 0, 8)
	if uint32(*t)&inotify.IN_ACCESS > 0 {
		parts = append(parts, "access")
	}
	if uint32(*t)&inotify.IN_ALL_EVENTS > 0 {
		parts = append(parts, "all")
	}
	if uint32(*t)&inotify.IN_ATTRIB > 0 {
		parts = append(parts, "attrib")
	}
	if uint32(*t)&inotify.IN_CLOSE > 0 {
		parts = append(parts, "close")
	}
	if uint32(*t)&inotify.IN_CLOSE_NOWRITE > 0 {
		parts = append(parts, "close-nowrite")
	}
	if uint32(*t)&inotify.IN_CLOSE_WRITE > 0 {
		parts = append(parts, "close-write")
	}
	if uint32(*t)&inotify.IN_CREATE > 0 {
		parts = append(parts, "create")
	}
	if uint32(*t)&inotify.IN_DELETE > 0 {
		parts = append(parts, "delete")
	}
	if uint32(*t)&inotify.IN_DELETE_SELF > 0 {
		parts = append(parts, "delete-self")
	}
	if uint32(*t)&inotify.IN_MODIFY > 0 {
		parts = append(parts, "modify")
	}
	if uint32(*t)&inotify.IN_MOVE > 0 {
		parts = append(parts, "move")
	}
	if uint32(*t)&inotify.IN_MOVED_FROM > 0 {
		parts = append(parts, "moved-from")
	}
	if uint32(*t)&inotify.IN_MOVED_TO > 0 {
		parts = append(parts, "moved-to")
	}
	if uint32(*t)&inotify.IN_MOVE_SELF > 0 {
		parts = append(parts, "move-self")
	}
	if uint32(*t)&inotify.IN_OPEN > 0 {
		parts = append(parts, "open")
	}
	if uint32(*t)&inotify.IN_ISDIR > 0 {
		parts = append(parts, "isdir")
	}
	if uint32(*t)&inotify.IN_IGNORED > 0 {
		parts = append(parts, "ignored")
	}
	if uint32(*t)&inotify.IN_Q_OVERFLOW > 0 {
		parts = append(parts, "q-overflow")
	}
	if uint32(*t)&inotify.IN_UNMOUNT > 0 {
		parts = append(parts, "unmount")
	}
	return strings.Join(parts, ",")
}

func (t *trackMask) Set(v string) error {
	parts := strings.Split(v, ",")
	for _, part := range parts {
		switch part {
		case "access":
			*t |= trackMask(inotify.IN_ACCESS)
		case "all":
			*t |= trackMask(inotify.IN_ALL_EVENTS)
		case "attrib":
			*t |= trackMask(inotify.IN_ATTRIB)
		case "close":
			*t |= trackMask(inotify.IN_CLOSE)
		case "close-nowrite":
			*t |= trackMask(inotify.IN_CLOSE_NOWRITE)
		case "close-write":
			*t |= trackMask(inotify.IN_CLOSE_WRITE)
		case "create":
			*t |= trackMask(inotify.IN_CREATE)
		case "delete":
			*t |= trackMask(inotify.IN_DELETE)
		case "delete-self":
			*t |= trackMask(inotify.IN_DELETE_SELF)
		case "modify":
			*t |= trackMask(inotify.IN_MODIFY)
		case "move":
			*t |= trackMask(inotify.IN_MOVE)
		case "moved-from":
			*t |= trackMask(inotify.IN_MOVED_FROM)
		case "moved-to":
			*t |= trackMask(inotify.IN_MOVED_TO)
		case "move-self":
			*t |= trackMask(inotify.IN_MOVE_SELF)
		case "open":
			*t |= trackMask(inotify.IN_OPEN)
		case "isdir":
			*t |= trackMask(inotify.IN_ISDIR)
		case "ignored":
			*t |= trackMask(inotify.IN_IGNORED)
		case "q-overflow":
			*t |= trackMask(inotify.IN_Q_OVERFLOW)
		case "unmount":
			*t |= trackMask(inotify.IN_UNMOUNT)
		default:
			return fmt.Errorf("unrecognized event type: %s", part)
		}
	}
	return nil
}
