package view

import (
	"time"

	"github.com/fhs/gompd/mpd"
)

type MainView struct {
	Artist   string
	Song     string
	Playlist []mpd.Attrs
	Time     time.Time
}
