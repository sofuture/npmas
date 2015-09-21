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
		c.Ui.SetStatus(&view.MainView{
			Artist:   "foo",
			Song:     "bar",
			Playlist: []string{"foo", "bar"},
			Time:     time.Now(),
		})
		time.Sleep(500 * time.Millisecond)
	}
}
