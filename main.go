package main

import (
	"log"
	"os"
)

func main() {
	log.Println("top -pid", os.Getpid())
	// gopp.PauseAk() // 到这儿，内存24M

	app := NewQGuiApplication(1, []string{"./heh.exe"}, 0)
	// gopp.PauseAk() // 到这儿，内存26M
	ape := NewQQmlApplicationEngine(nil)
	// gopp.PauseAk() // 到这儿，内存27M
	ape.Load("hh.qml")
	log.Println("top -pid", os.Getpid())

	// gopp.PauseAk() // 到这儿，内存28M
	log.Println("app.Exec ...", "top -pid", os.Getpid())
	app.Exec() // 到这儿，内存32M

}
