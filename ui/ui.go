package ui

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"

	"github.com/sofuture/npmas/view"
)

type Ui struct {
	Title string

	gui                      *gocui.Gui
	header, footer, playlist *gocui.View
	shutdown                 chan string
}

func (u *Ui) getLayout() func(*gocui.Gui) error {
	return func(g *gocui.Gui) error {

		var err error

		maxX, maxY := g.Size()

		topHeight := 2
		bottomHeight := 3

		if u.header, err = g.SetView("header", -1, -1, maxX, topHeight); err != nil {
			if err != gocui.ErrorUnkView {
				return err
			}
			u.header.BgColor = gocui.ColorRed
			u.header.Frame = false
		}

		if u.playlist, err = g.SetView("playlist", -1, topHeight-1, maxX, maxY-bottomHeight); err != nil {
			if err != gocui.ErrorUnkView {
				return err
			}
			u.playlist.BgColor = gocui.ColorBlue
			u.playlist.Frame = false
		}

		if u.footer, err = g.SetView("footer", -1, maxY-(1+bottomHeight), maxX, maxY); err != nil {
			if err != gocui.ErrorUnkView {
				return err
			}
			u.footer.Autoscroll = true
			u.footer.BgColor = gocui.ColorGreen
			u.footer.Frame = false
		}
		return nil
	}
}

func (u *Ui) Run() {
	var err error
	u.gui = gocui.NewGui()
	if err = u.gui.Init(); err != nil {
		log.Panicln(err)
	}
	defer u.gui.Close()

	u.gui.SetLayout(u.getLayout())

	quit := func(g *gocui.Gui, v *gocui.View) error {
		return gocui.Quit
	}

	if err := u.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	err = u.gui.MainLoop()
	if err != nil && err != gocui.Quit {
		log.Panicln(err)
	}
}

func (u *Ui) SetStatus(status *view.MainView) {
	if u.header != nil {
		u.header.Clear()
		fmt.Fprintf(u.header, "%s - %s", status.Artist, status.Song)
	}

	if u.playlist != nil {
		u.playlist.Clear()
		for _, k := range status.Playlist {
			fmt.Fprintf(u.playlist, "%s\n", k)
		}
	}

	if u.footer != nil {
		fmt.Fprintln(u.footer, status.Time)
	}

	u.gui.Flush()
}
