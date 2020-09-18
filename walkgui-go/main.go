package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strings"
	"time"
)

func main() {
	var inTE, outTE *walk.TextEdit

	go func() {
		time.Sleep(time.Second * 5)
		for {
			outTE.AppendText("hello")
			time.Sleep(time.Second)
		}
	}()

	ret, err := MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()

	fmt.Println(ret, err)
}
