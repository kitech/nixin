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

	// mw := qtwidgets.NewQMainWindow(nil, 0)
	// mw.Show()
	// mw.Hide()
	// mw.Dtor()

	mw := NewMainWindow()
	mw.Show()

	qmlmain()

	// gopp.PauseAk() // 到这儿，内存28M
	log.Println("app.Exec ...", "top -pid", os.Getpid(), "lsof -p", os.Getpid())
	app.Exec() // 到这儿，内存32M
}
