package main

import (
	"log"
	"os"

	_ "github.com/kitech/nixcmd"
	"github.com/qtui/qtsyms"
	"github.com/qtui/qtwidgets"
)

func main() {
	log.Println("top -pid", os.Getpid(), "lsof -p", os.Getpid())
	// gopp.PauseAk() // 到这儿，内存24M
	qtsyms.LoadAllQtSymbols() // 这个初始化现在是必须的

	app := qtwidgets.NewQApplication([]string{"./nixui.exe"}, 0)
	// gopp.PauseAk() // 到这儿，内存26M
	// ape := NewQQmlApplicationEngine(nil)
	// gopp.PauseAk() // 到这儿，内存27M
	// ape.Load("hh.qml")
	// log.Println("top -pid", os.Getpid())

	// btn := qtwidgets.NewQPushButton()
	// qtrt.Connect(btn, "clicked(bool)", func(b bool) {
	// 	log.Println("works???", b)
	// })
	// btn.SetFlat(true)
	// btn.SetText("hehhe")
	// btn.Show()

	// btn.Size()

	// btn.Dtor()

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
