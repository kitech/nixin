package main

import (
	"log"

	"github.com/qtui/qtrt"
	"github.com/qtui/qtwidgets"
)

type MainWindow struct {
	*qtwidgets.QMainWindow

	split1 *qtwidgets.QSplitter
	split2 *qtwidgets.QSplitter
	split3 *qtwidgets.QSplitter

	lftw1   *qtwidgets.QWidget
	lovbox1 *qtwidgets.QVBoxLayout // left nav
	lohbox1 *qtwidgets.QVBoxLayout // right top nav

	stkw1 *qtwidgets.QStackedWidget

	tabw1 *qtwidgets.QTabWidget
	tabw2 *qtwidgets.QTabWidget

	lftbtnmenus []*qtwidgets.QPushButton
}

func NewMainWindow() *MainWindow {
	me := &MainWindow{}
	me.setupui()
	return me
}
func (me *MainWindow) setupui() {
	me.QMainWindow = qtwidgets.NewQMainWindow(nil, 0)
	me.split1 = qtwidgets.NewQSplitter(me.QMainWindow.QWidget)
	me.SetCentralWidget(me.split1.QWidget)

	{
		btn := qtwidgets.NewQPushButton("hehee111")
		qtrt.Connect(btn, "clicked(bool)", func(b bool) {
			log.Println("works???", b)
		})
		// btn.SetFlat(true)
		// btn.SetText("hehhe111")
		me.split1.AddWidget(btn.QWidget)
	}
	{
		btn := qtwidgets.NewQPushButtonz0()
		qtrt.Connect(btn, "clicked(bool)", func(b bool) {
			log.Println("works???222", b)
		})
		// btn.SetFlat(true)
		btn.SetText("hehhe222")
		me.split1.AddWidget(btn.QWidget)
	}

	btnmenudats := []any{
		"hehehhe333", func(b bool) { log.Println("hehehhe333") },
		"hehehhe444", func(b bool) { log.Println("hehehhe444") },
		"hehehhe666", func(b bool) { log.Println("hehehhe555") },
		"hehehhe777", func(b bool) { log.Println("hehehhe666") },
		"hehehhe888", func(b bool) { log.Println("hehehhe777") },
		"hehehhe999", func(b bool) { log.Println("hehehhe888") },
		"hehehheaaa", func(b bool) { log.Println("hehehhe999") },
		"hehehhebbb", func(b bool) { log.Println("hehehhebbb") },
		"hehehheccc", func(b bool) { log.Println("hehehheccc") },
		"hehehheddd", func(b bool) { log.Println("hehehheddd") },
	}

	me.lovbox1 = qtwidgets.NewQVBoxLayout(me.split1.QWidget)
	for i := 0; i < len(btnmenudats); i += 2 {
		name := btnmenudats[i].(string)
		fn := btnmenudats[i+1].(func(bool))

		btn := qtwidgets.NewQPushButtonz0()
		btn.SetText(name)
		qtrt.Connect(btn, "clicked(bool)", fn)

		me.lftbtnmenus = append(me.lftbtnmenus, btn)
		me.lovbox1.AddWidget(btn.QWidget)
	}

	me.lftw1 = qtwidgets.NewQWidget(nil)
	me.lftw1.SetLayout(me.lovbox1.QLayout)
	me.split1.AddWidget(me.lftw1)
}
