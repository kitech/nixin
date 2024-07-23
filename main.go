package main

import (
	"log"
	"os"

	"github.com/qtui/qtcore"
	"github.com/qtui/qtsyms"
	"github.com/qtui/qtwidgets"
)

func main() {
	log.Println("top -pid", os.Getpid(), "lsof -p", os.Getpid())
	// gopp.PauseAk() // 到这儿，内存24M
	qtsyms.LoadAllQtSymbols() // 这个初始化现在是必须的

	app := qtwidgets.NewQApplication([]string{"./heh.exe"}, 0)
	// gopp.PauseAk() // 到这儿，内存26M
	// ape := NewQQmlApplicationEngine(nil)
	// gopp.PauseAk() // 到这儿，内存27M
	// ape.Load("hh.qml")
	// log.Println("top -pid", os.Getpid())

	btn := qtwidgets.NewQPushButton()
	btn.SetFlat(true)
	btn.SetText("hehhe")
	btn.Show()

	qstr := qtcore.QString_FromUtf8("hehhee")
	log.Println(qstr.Length())
	qstr.Dtor()

	// gopp.PauseAk() // 到这儿，内存28M
	log.Println("app.Exec ...", "top -pid", os.Getpid(), "lsof -p", os.Getpid())
	app.Exec() // 到这儿，内存32M
}
