package view

import (
	"time"
)

type MainView struct {
	Artist   string
	Song     string
	Playlist []string
	Time     time.Time
}
