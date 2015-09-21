package client

import (
	"log"
	"time"

	"github.com/fhs/gompd/mpd"

	"github.com/sofuture/npmas/ui"
	"github.com/sofuture/npmas/view"
)

type Client struct {
	Conn *mpd.Client
	Ui   *ui.Ui
}

const updateHz = 1

func Run(server string) int {
	conn, err := mpd.Dial("tcp", server)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	defer conn.Close()

	u := &ui.Ui{Title: "nmpas"}
	c := &Client{Conn: conn, Ui: u}
	go c.Run()
	u.Run()

	return 0
}

func (c *Client) Run() {
	for {
		current, err := c.Conn.CurrentSong()
		if err != nil {
			return
		}

		plinfo, err := c.Conn.PlaylistInfo(-1, -1)
		if err != nil {
			return
		}

		c.Ui.SetStatus(&view.MainView{
			Artist:   current["artist"],
			Song:     current["song"],
			Playlist: plinfo,
			Time:     time.Now(),
		})
		time.Sleep(time.Second / updateHz)
	}
}

func (c *Client) TogglePlay() {
	c.Conn.Pause(false)
}
