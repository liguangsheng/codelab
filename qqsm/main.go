package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/qianniancn/go-dmsoft"
	"os"
	"strings"
	"time"
)

func main() {
	var inTE, outTE *walk.TextEdit

	mainWindow := MainWindow{
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
	}

	// 如果是x64编译或运行 需要设置export GOARCH=386
	// 使用免注册方式注册大漠插件，也可以使用命令行注册
	// 由于大漠是付费插件，某些功能可以免费使用，但是后台高级功能需要付费
	dir, err := os.Getwd()
	fmt.Println(dir, err)
	// 这里是免注册方式
	dmsoft.SetDllPathW(dir+"\\dm.dll", 0)

	dm, err := dmsoft.New()
	if err != nil {
		panic(err)
	}
	defer dm.Release() //必须释放
	fmt.Printf("当前插件版本:%s\n", dm.Ver())

	go func() {
		time.Sleep(time.Second * 5)
		outTE.SetText(fmt.Sprintf("当前插件版本:%s\n", dm.Ver()))
	}()

	mainWindow.Run()
}

func InitDM() {
	// 如果是x64编译或运行 需要设置export GOARCH=386
	// 使用免注册方式注册大漠插件，也可以使用命令行注册
	// 由于大漠是付费插件，某些功能可以免费使用，但是后台高级功能需要付费
	dir, _ := os.Getwd()
	// 这里是免注册方式
	dmsoft.SetDllPathW(dir+"\\bin\\bin.dll", 0)

	dm, err := dmsoft.New()
	if err != nil {
		panic(err)
	}
	defer dm.Release() //必须释放

	fmt.Printf("当前插件版本:%s\n", dm.Ver())
}
