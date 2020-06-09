package main

import (
	"github.com/leaanthony/mewn"
	"github.com/lrtz-v/Tutorial/gui/cpustats/pks/sys"
	"github.com/wailsapp/wails"
	// "gui/cpustats/pkg/sys"
)

func basic() string {
	return "Hello World!"
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	stats := &sys.Stats{}

	app := wails.CreateApp(&wails.AppConfig{
		Width:  512,
		Height: 512,
		Title:  "CPU Usage",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(stats)
	app.Run()
}
