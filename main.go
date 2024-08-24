package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/kitech/nixcmd"
	"github.com/qtui/qtcore"
	"github.com/qtui/qtqml"
	"github.com/qtui/qtsyms"
	"github.com/qtui/qtwidgets"
)

func main() {
	flag.Parse()
	typ := flag.Arg(0)
	log.Println(typ)

	if typ == "qml" {
		mainqml()
	} else {
		mainwgt()
	}
}

func mainqml() {
	log.Println("top -pid", os.Getpid(), "lsof -p", os.Getpid())
	// gopp.PauseAk() // 到这儿，内存24M
	qtsyms.LoadAllQtSymbols() // 这个初始化现在是必须的

	app := qtcore.NewQGuiApplication([]string{"./nixui.exe"}, 0)

	qmlview := qtqml.NewQQuickView()
	qmlview.Show()
	qmlview.SetSource("qmlview.qml")

	vrobj := qmlview.ContentItem()
	log.Println(vrobj.Dbgstr())

	// vlo := qtqml.NewQQuickColumnLayout(vrobj)
	// _ = vlo
	toplox := vrobj.FindChild("toplo")
	toplo := qtqml.QQuickColumnLayoutFromptr(toplox.GetCthis())

	for i := 0; i < 30; i++ {
		if true {
			break
		}
		lb := qtqml.NewQQuickLabel(toplo)
		lb.SetText(fmt.Sprintf("heowiwf %d", i))
		lb.SetProperty("color", "white")
		ta := qtqml.NewQQuickTextArea(toplo)
		ta.SetProperty("color", "white")
		ta.SetProperty("text", "whitealiefieaf瀑 fewer iaefjoeff粗盐")
		ta.SetWidth(100)
	}

	// qmlmain()

	log.Println("top -pid", os.Getpid(), "lsof -p", os.Getpid())
	app.Exec()
}

func mainwgt() {
	log.Println("top -pid", os.Getpid(), "lsof -p", os.Getpid())
	// gopp.PauseAk() // 到这儿，内存24M
	qtsyms.LoadAllQtSymbols() // 这个初始化现在是必须的

	app := qtwidgets.NewQApplication([]string{"./nixui.exe"}, 0)
	// gopp.PauseAk() // 到这儿，内存26M

	// mw := qtwidgets.NewQMainWindow(nil, 0)
	// mw.Show()
	// mw.Hide()
	// mw.Dtor()

	mw := NewMainWindow()
	mw.Show()

	// gopp.PauseAk() // 到这儿，内存28M
	log.Println("app.Exec ...", "top -pid", os.Getpid(), "lsof -p", os.Getpid())
	app.Exec() // 到这儿，内存32M
}
